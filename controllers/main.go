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
	strs := Obj.GetConfig()
	c.Data["url_test"] = strs[0]
	c.Data["url_qa"] = strs[1]
	c.Data["username"] = strs[2]
	c.Data["password"] = strs[3]
	c.TplName = "index.tpl"
}
