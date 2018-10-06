package controllers

import (
	"quickstart/models"
	"fmt"
	"strconv"
)
type ShowcommentsController struct {
	BaseController
}
func (this *ShowcommentsController) Post(){
	comments_id := this.GetString("comments_id")
	fmt.Println("接收到的comments_id",comments_id)
	int_id,_:=strconv.Atoi(comments_id)

    comments:=models.GetRoomCommentRes(int_id)
    this.Data["json"]=comments
    this.ServeJSON()
}
