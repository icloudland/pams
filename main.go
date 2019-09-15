package main

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	beego.Run()

}

type MainController struct {
	beego.Controller
}

func init() {
	beego.Router("/", &MainController{}, "*:Index")
	beego.Router("/index.html", &MainController{}, "*:Index")
	beego.Router("/login.html", &MainController{}, "*:Index")
	beego.Router("/admin.html", &MainController{}, "*:Index")
	beego.Router("/active.html", &MainController{}, "*:Index")
	beego.Router("/info.html", &MainController{}, "*:Index")
	beego.Router("/table.html", &MainController{}, "*:Index")
	beego.Router("/welcome.html", &MainController{}, "*:Index")
	beego.Router("/form.html", &MainController{}, "*:Index")

	beego.Router("/api/init.json", &ApiController{}, "*:InitData")
	beego.Router("/api/table.json", &ApiController{}, "*:InitData")
}

func (c *MainController) Index() {

	tplName := strings.TrimPrefix(c.Ctx.Request.URL.Path, "/")

	if tplName == "" {
		tplName = "index.html"
	}

	c.TplName = tplName
}

type ApiController struct {
	beego.Controller
}

func (c *ApiController) InitData() {

	fileName := strings.TrimPrefix(c.Ctx.Request.URL.Path, "/")

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	c.Ctx.Output.Body(data)
	c.StopRun()
}
