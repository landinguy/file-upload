package controllers

import (
	"file-upload/util"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) CreateTable() {
	util.CreateTable()
	c.Ctx.WriteString("表创建成功")
}
