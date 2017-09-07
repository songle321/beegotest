package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit")
	fmt.Println(isExit)
	if isExit == "true" {
		fmt.Println("remove cookie")
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
	}
	
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	pwd := this.GetString("pwd")
	fmt.Println("username:", username, ",pwd:", pwd)
	if beego.AppConfig.String("username") == username && beego.AppConfig.String("pwd") == pwd {
		maxage := 0
		fmt.Println("login")
		this.Ctx.SetCookie("username", username, maxage, "/")
		this.Ctx.SetCookie("pwd", pwd, maxage, "/")
		this.Ctx.Redirect(301, "/")
		return
	} else {
		this.Ctx.Redirect(301, "/login")
	}
	return
}
func checkAccount(Ctx *context.Context) bool {
	_, err := Ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := Ctx.GetCookie("username")
	pwd := Ctx.GetCookie("pwd")
	if username == beego.AppConfig.String("username") && pwd == beego.AppConfig.String("pwd") {
		return true
	} else {
		return false
	}
}
