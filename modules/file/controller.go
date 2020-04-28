package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"

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

// @router /file/line/:file/?:line([0-9]*) [get]
func (c *FileController) Line() {
	filename := c.ParamStr("file")
	linenum := c.ParamInt("line")

	ret := make(map[int]string)

	file, err := os.Open(dir + filename)
	c.ReturnErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := 0; i < linenum; i++ {
		_, err := reader.ReadString('\n')
		if io.EOF == err {
			break
		}
	}

	reg := regexp.MustCompile("\x1b[^mK]*[mK]")
	for i := linenum; i < linenum+1000; i++ {
		line, err := reader.ReadString('\n')

		line = reg.ReplaceAllString(line, "")

		if io.EOF == err {
			break
		}

		if i >= linenum {
			ret[i] = line[:len(line)-1]
		}
	}

	c.Return(ret)
}
