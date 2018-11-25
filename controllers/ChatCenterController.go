package controllers

import "quickstart/models"
type ChatCenterController struct {
	BaseController
}
func (this *ChatCenterController)Get(){
    CurrentId := this.GetSession("user_id")
    if CurrentId == nil{
    	this.Redirect("/login",302)
	}
    FromDatas,_:=models.GetALLChatFromInfo(CurrentId.(int))
    ToDatas,_:=models.GetALLChatToInfo(CurrentId.(int))
	this.Data["FromDatas"]=FromDatas
	this.Data["ToDatas"] =ToDatas
	this.Data["CurrentId"] =CurrentId
	this.Layout = "layout/layout.html"
	this.TplName = "chat/center.html"
}
