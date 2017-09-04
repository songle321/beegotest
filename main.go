package main

import (
	_ "blog/routers"

	_ "blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	//orm.RegisterDataBase("default", "mysql", "root:111111@/blog?charset=utf8", 30, 30) //注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:111111@/blog?charset=utf8", 10, 10)

}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	// o := orm.NewOrm()
	// o.Using("default")
	beego.SetStaticPath("static", "static")
	beego.Run()
}
