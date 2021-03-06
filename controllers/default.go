package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["TrueCond"] = true
	c.TplName = "home.html"
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
