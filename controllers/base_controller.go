package controllers

import (
	"encoding/json"
	"os"

	"github.com/YouDad/localhost/log"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Error404() {
	c.TplName = "index.html"
}

func (c *BaseController) Get() {
	c.TplName = "index.html"
}

type SimpleJSONResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *BaseController) ParseParameter(data interface{}) {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, data)
	if err != nil {
		log.Infof("%s\n", c.Ctx.Input.RequestBody)
		log.Err(err)
		c.ReturnErr(err)
	}
	log.SetCallerLevel(1)
	log.Infoln(data)
	log.SetCallerLevel(0)
}

func (c *BaseController) GetCookie(key string) string {
	return c.Ctx.GetCookie(key)
}

func (c *BaseController) GetToken() string {
	return c.GetCookie("token")
}

func (c *BaseController) ReturnJson(data interface{}) {
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) ReturnErr(err error) {
	if err != nil {
		c.ReturnJson(SimpleJSONResult{
			Code:    1,
			Message: err.Error(),
		})
	}
}

func (c *BaseController) ReturnMessage(message string) {
	c.ReturnJson(SimpleJSONResult{0, message})
}

func (c *BaseController) Return(data interface{}) {
	c.ReturnJson(struct {
		SimpleJSONResult
		Data interface{} `json:"data"`
	}{SimpleJSONResult{Code: 0}, data})
}

func (c *BaseController) ReturnErrNoPermission() {
	c.ReturnJson(SimpleJSONResult{401, "无权访问"})
}

func (c *BaseController) Param(key string) string {
	return c.Ctx.Input.Param(key)
}

func ReadJsonFromFile(jsonPath string, dataRef interface{}) error {
	file, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	log.Err(err)

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	log.Err(err)

	return json.Unmarshal(data, &dataRef)
}
