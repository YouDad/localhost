package test

import (
	"io/ioutil"

	"github.com/YouDad/localhost/controllers"
	"github.com/astaxie/beego"
)

type FileController struct {
	controllers.BaseController
}

func init() {
	beego.Router("/file/list", new(FileController), "get:List")
}

// @router /file/list [get]
func (c *FileController) List() {
	dir, err := ioutil.ReadDir("/home/manjaro/Log/blockchain")
	c.ReturnErr(err)

	var ret []interface{}

	for i := len(dir) - 1; i >= 0; i-- {
		if dir[i].IsDir() {
			continue
		}
		ret = append(ret, struct {
			Name string
			Size int64
		}{dir[i].Name(), dir[i].Size()})
	}

	c.Return(ret)
}
