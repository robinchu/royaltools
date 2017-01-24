package controllers

import (
	_ "encoding/json"
	"github.com/astaxie/beego"
	"royaltools/service"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	Obj := service.ValidateFlashService{4, 9}
	c.Data["area"] = Obj.Area()
	c.Data["config"] = Obj.GetConfig()
	c.TplName = "index.tpl"
}
