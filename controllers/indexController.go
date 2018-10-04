package controllers

import (

	_"github.com/astaxie/beego/orm"

	_"quickstart/models"
	//"fmt"
)

type IndexController struct {
	BaseController
}
func (this *IndexController) Get(){
	this.Layout = "layout/layout.html"

	this.TplName = "main/index.html"
}