package controllers

import (

	"github.com/astaxie/beego/orm"

	"quickstart/models"
	_"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	c.Data["Email"] = "astaxie@gmail.com"
	orm.Debug = true
	ss := models.GetAllRooms()

	c.Layout = "layout/layout.html"
	c.Data["Datas"]=ss

	c.TplName = "main/roomall.html"
}
