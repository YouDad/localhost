package modules

import (
	"github.com/YouDad/localhost/modules/file"
	"github.com/YouDad/localhost/modules/test"
	"github.com/astaxie/beego"
)

func Init() {
	beego.Include(new(file.FileController))
	beego.Include(new(test.TestController))
}
