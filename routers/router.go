package routers

import (
	"github.com/YouDad/localhost/controllers"
	"github.com/YouDad/localhost/modules"
	"github.com/astaxie/beego"
)

func init() {
	modules.Init()

	beego.ErrorController(&controllers.BaseController{})
	beego.SetStaticPath("/", "views/")
}
