package file

import (
	"io/ioutil"

	"github.com/YouDad/localhost/controllers"
)

type FileController struct {
	controllers.BaseController
}

const dir = "/home/manjaro/Log/blockchain/"

// @router /file/list [get]
func (c *FileController) List() {
	dir, err := ioutil.ReadDir(dir)
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
