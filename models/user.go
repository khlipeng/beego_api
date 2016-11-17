package models

import (
	"github.com/astaxie/beego/orm"
	// "log"
	// "fmt"
	"time"
)

type User struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
	Nickname string    `json:"nickname" orm:"column(nickname);unique;size(40);"`
	Password string    `json:"-" orm:"column(password);size(40)"`
	Created  time.Time `json:"create_at" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return TableName("user")
}
func init() {
	orm.RegisterModel(new(User))
}
func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// 检测手机号是否注册
func CheckUserPhone(phone string) bool {
	exist := Users().Filter("phone", phone).Exist()
	return exist
}

// 检测用户昵称是否存在
func CheckUserNickname(nickname string) bool {
	exist := Users().Filter("nickname", nickname).Exist()
	return exist
}

//创建用户
func CreateUser(user User) User {
	o := orm.NewOrm()
	o.Insert(&user)
	return user
}

//检测手机和昵称是否注册
func CheckUserPhoneOrNickname(phone string, nickname string) bool {
	cond := orm.NewCondition()
	count, _ := Users().SetCond(cond.And("phone", phone).Or("nickname", nickname)).Count()
	if count <= int64(0) {
		return false
	}
	return true
}
func CheckUserAuth(nickname string, password string) (User, bool) {
	o := orm.NewOrm()
	user := User{
		Nickname: nickname,
		Password: password,
	}
	err := o.Read(&user, "Nickname", "Password")
	if err != nil {
		return user, false
	}
	return user, true
}

// User database CRUD methods include Insert, Read, Update and Delete
func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}
