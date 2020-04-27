package routers

import (
	"github.com/YouDad/localhost/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.BaseController{})
	beego.SetStaticPath("/", "views/")
}
