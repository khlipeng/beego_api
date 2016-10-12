package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = Response{
		Errcode: 404,
		Errmsg:  "Not Found",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = Response{
		Errcode: 401,
		Errmsg:  "Permission denied",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = Response{
		Errcode: 403,
		Errmsg:  "Forbidden",
	}
	c.ServeJSON()
}
