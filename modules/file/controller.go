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

	var unmatch []interface{}
	var match []interface{}

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

	reg, err := regexp.Compile("\x1b[^mK]*[mK]")
	c.ReturnErr(err)

	extract, err := regexp.Compile(`(.*/.*/.*) (.*:.*:.*\..*)\[(.*)\]\[(.*)\]\[(.*)\]: { (.* .*) } (.*)`)
	c.ReturnErr(err)

	for i := linenum; i < linenum+10000; i++ {
		line, err := reader.ReadString('\n')

		line = reg.ReplaceAllString(line, "")
		line = line[:len(line)-1]
		submatch := extract.FindStringSubmatch(line)

		if io.EOF == err {
			break
		}

		if len(submatch) < 8 {
			unmatch = append(unmatch, struct {
				Line    int    `json:"line"`
				Content string `json:"content"`
			}{i, line})
		} else {
			match = append(match, struct {
				Line    int    `json:"line"`
				Date    string `json:"date"`
				Time    string `json:"time"`
				Port    string `json:"port"`
				Routine string `json:"routine"`
				Level   string `json:"level"`
				Source  string `json:"source"`
				Content string `json:"content"`
			}{i, submatch[1], submatch[2], submatch[3], submatch[4],
				submatch[5], submatch[6], submatch[7]})
		}
	}

	c.Return(struct {
		Unmatch interface{} `json:"unmatch"`
		Match   interface{} `json:"match"`
	}{unmatch, match})
}
