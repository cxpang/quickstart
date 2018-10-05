package controllers

import (
	"fmt"
	"quickstart/models"
	"strings"

)

type RoomdetailController struct {
	BaseController
}
func (this *RoomdetailController) Get(){
    room_id:=this.Ctx.Input.Param(":roomid")
    room_detail:=models.GetRoomDetail(room_id)
    room_image:=room_detail.Room_images
    image_array:=strings.Split(room_image, ",")
	fmt.Println("获取的房间",image_array)
    this.Data["roomdetail"]=room_detail
	this.Data["image0"] = image_array[0]
	this.Data["image1"] = image_array[1]
	this.Data["image2"] = image_array[2]
	this.Layout = "layout/layout.html"
	this.TplName = "main/roomdetail.html"
}