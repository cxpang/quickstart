package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"

	_"quickstart/models"
	//"fmt"
)

type ShowController struct {
	beego.Controller
}

func (c *ShowController) Get() {
	c.TplName = "show.html"
}
