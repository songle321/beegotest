package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Category), new(Topic))                                       //注册 model                                                 //注册 model
}

/*文章的目录*/
type Category struct {
	ID              int64
	Title           string
	Create          time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

/*Topic:文章*/
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Comtent         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}
