package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"quickstart/models"
	//"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Email"] = "astaxie@gmail.com"
	orm.Debug = true
	ss := models.GetAllRooms()
	c.Layout = "layout/layout.html"
	c.Data["Datas"]=ss


	user:=c.GetSession("user_name")
	if user==nil {
		c.Data["IsLogin"]=1
		c.Data["User_name"]="未登录"
	}else {
		c.Data["IsLogin"]=2
		c.Data["User_name"]=user
		}

	c.TplName = "show.html"
}
