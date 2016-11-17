package controllers

import (
	// "crypto"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/khlipeng/beego_api/models"
	"github.com/khlipeng/beego_api/utils"
	// "log"
	"strings"
	"time"
)

var (
	ErrPhoneIsRegis     = ErrResponse{422001, "手机用户已经注册"}
	ErrNicknameIsRegis  = ErrResponse{422002, "用户名已经被注册"}
	ErrNicknameOrPasswd = ErrResponse{422003, "账号或密码错误。"}
)

type UserController struct {
	BaseController
}
type LoginToken struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// @Title 注册新用户
// @Description 用户注册
// @Param	phone		formData 	string	true 		"用户手机号"
// @Param	nickname	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码(需要前端 Md5 后传输)"
// @Success 200 {object}
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (this *UserController) Registered() {
	phone := this.GetString("phone")
	nickname := this.GetString("nickname")
	password := this.GetString("password")

	valid := validation.Validation{}
	//表单验证
	valid.Required(phone, "phone").Message("手机必填")
	valid.Required(nickname, "nickname").Message("用户昵称必填")
	valid.Required(password, "password").Message("密码必填")
	valid.Mobile(phone, "phone").Message("手机号码不正确")
	valid.MinSize(nickname, 2, "nickname").Message("用户名最小长度为 2")
	valid.MaxSize(nickname, 40, "nickname").Message("用户名最大长度为 40")
	valid.Length(password, 32, "password").Message("密码格式不对")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	if models.CheckUserPhone(phone) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrPhoneIsRegis
		this.ServeJSON()
		return
	}
	if models.CheckUserNickname(nickname) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrNicknameIsRegis
		this.ServeJSON()
		return
	}

	user := models.User{
		Phone:    phone,
		Nickname: nickname,
		Password: password,
	}
	this.Data["json"] = Response{0, "success.", models.CreateUser(user)}
	this.ServeJSON()

}

// @Title 登录
// @Description 账号登录
// @Success 200 {object}
// @Failure 404 no enough input
// @Failure 401 No Admin
// @router /login [post]
func (this *UserController) Login() {
	nickname := this.GetString("nickname")
	password := this.GetString("password")

	user, ok := models.CheckUserAuth(nickname, password)
	if !ok {
		this.Data["json"] = ErrNicknameOrPasswd
		this.ServeJSON()
		return
	}

	et := utils.EasyToken{
		Username: user.Nickname,
		Uid:      user.Id,
		Expires:  time.Now().Unix() + 3600,
	}

	token, err := et.GetToken()
	if token == "" || err != nil {
		this.Data["json"] = ErrResponse{-0, err}
	} else {
		this.Data["json"] = Response{0, "success.", LoginToken{user, token}}
	}

	this.ServeJSON()
}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {object}
// @Failure 404 no enough input
// @Failure 401 No Admin
// @router /auth [get]
func (this *UserController) Auth() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(authtoken)
	if !valido {
		this.Data["json"] = ErrResponse{-1, fmt.Sprintf("%s", err)}
		this.ServeJSON()
		return
	}

	this.Data["json"] = Response{0, "success.", "is login"}
	this.ServeJSON()
}
