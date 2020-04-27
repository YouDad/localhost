package test

import (
	"github.com/YouDad/localhost/controllers"
	"github.com/astaxie/beego"
)

type TestController struct {
	controllers.BaseController
}

func init() {
	beego.Router("/test", new(TestController), "get:Test")
}

// @router /test [get]
func (c *TestController) Test() {
}
