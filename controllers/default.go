package controllers

// Operations about object
type DefaultController struct {
	BaseController
}

// @Title 欢迎信息
// @Description API 欢迎信息
// @Success 200 {object}
// @router / [any]
func (o *DefaultController) GetAll() {
	o.Data["json"] = Response{0, "success.", "API 1.0"}
	o.ServeJSON()
}
