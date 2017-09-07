package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Category), new(Topic)) //注册 model                                                 //注册 model
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

func AddCategory(name string) error {
	db := orm.NewOrm()
	cate := &Category{Title: name, Create: time.Now(), Views: 0, TopicTime: time.Now()}
	qs := db.QueryTable(&Category{})
	err := qs.Filter("Title", name).One(cate)
	if err == nil {

		return err
	}
	fmt.Println("insert")
	_, error := db.Insert(cate)
	if error != nil {
		return error
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	db := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := db.QueryTable(&Category{})
	_, err := qs.All(&cates)
	if err != nil {
		return nil, err
	}
	return cates, nil

}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	db := orm.NewOrm()
	cate := &Category{ID: cid}
	_, error := db.Delete(cate, "ID")
	return error
}
