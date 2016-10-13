package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/khlipeng/beego_api/models"
	"log"
)

type UserController struct {
	BaseController
}

// @Title 注册新用户
// @Description 用户注册
// @Param	phone		formData 	string	true		"用户手机号"
// @Param	nickname	formData 	string	true		"用户昵称"
// @Param	password	formData 	string	true		"密码(需要前端 Md5 后传输)"
// @Success 200 {object}
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (u *UserController) Registered() {
	phone := u.GetString("phone")
	nickname := u.GetString("nickname")
	password := u.GetString("password")

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
			u.Ctx.ResponseWriter.WriteHeader(403)
			u.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			u.ServeJSON()
			return
		}
	}
	if models.CheckUserPhone(phone) {
		u.Ctx.ResponseWriter.WriteHeader(422)
		u.Data["json"] = ErrResponse{422001, "手机用户已经注册"}
		u.ServeJSON()
		return
	}
	if models.CheckUserNickname(nickname) {
		u.Ctx.ResponseWriter.WriteHeader(422)
		u.Data["json"] = ErrResponse{422002, "用户名已经被注册"}
		u.ServeJSON()
		return
	}

	user := models.User{
		Phone:    phone,
		Nickname: nickname,
		Password: password,
	}
	u.Data["json"] = Response{0, "success.", models.CreateUser(user)}
	u.ServeJSON()

	log.Println("...")
}

// @Title 测试
// @Description 测试错误码
// @Success 200 {object}
// @Failure 404 no enough input
// @Failure 401 No Admin
// @router /test [get]
func (u *UserController) Test() {
	u.Data["json"] = Response{0, "success", "API 1.0"}
	u.Abort("404")
	u.ServeJSON()
}
