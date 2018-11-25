package controllers

import (
	"quickstart/models"
	"time"

)

type ChatController struct {
	BaseController
}
func (this *ChatController)Get(){
	userId,_:=this.GetInt(":touser")
	data,isbool:=models.GetChatUserInfo(userId)
	if isbool == false{
		this.Redirect("/",200)
	}
    time:=time.Now()
    CurrentId :=this.GetSession("user_id")
    if CurrentId == nil{
		this.Redirect("/login",302)
	}
    bool := models.CheckUserChatIsExist(CurrentId.(int),userId)
    if bool == false{
         models.SaveChatInfo(CurrentId.(int),userId)
	}
    this.Data["now"] = time
	this.Data["toUser"]=data
	this.Data["CurrentId"] =CurrentId
	this.Layout = "layout/layout.html"
	this.TplName = "chat/index.html"
}