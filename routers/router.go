package routers

import (
	"github.com/astaxie/beego"
	"github.com/khlipeng/beego_api/controllers"
)

// 使用注释路由
func init() {

	beego.Router("/", &controllers.DefaultController{}, "*:GetAll")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
