package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/khlipeng/beego_api/controllers:DefaultController"] = append(beego.GlobalControllerRouter["github.com/khlipeng/beego_api/controllers:DefaultController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"any"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/khlipeng/beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/khlipeng/beego_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Registered",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/khlipeng/beego_api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/khlipeng/beego_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
