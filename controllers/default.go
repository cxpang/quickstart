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
	//c.Data["Website"] = "pangchenxu"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	orm.Debug = true
	ss := models.GetAllRooms()
	c.Data["Datas"]=ss
	c.TplName = "header.html"
}
