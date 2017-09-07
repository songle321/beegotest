package controllers

import (
	"blog/models"
	"fmt"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsCatogory"] = true
	this.TplName = "category.html"

	opt := this.Input().Get("opt")
	switch opt {
	case "add":
		name := this.Input().Get("name")
		fmt.Println("opt:", opt, ";name:"+name)

		if len(name) == 0 {
			return
		}
		error := models.AddCategory(name)
		if error != nil {
			beego.Error(error)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		error := models.DelCategory(id)
		if error != nil {
			beego.Error(error)
		}
	}
	cates, error := models.GetAllCategories()
	if error != nil {
		beego.Error(error)
	}
	fmt.Println("end")
	this.Data["Categories"] = cates
}

func (this *CategoryController) Post() {
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
