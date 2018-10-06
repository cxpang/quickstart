package models

import (
	"github.com/astaxie/beego/orm"
   "fmt"
)
type RoomComments struct {
	Comments_id int  `orm:"pk;column(comments_id);"`
	Room_id int
	Comments_content string
	Created_at string
	Answer_count int
	User_id int
	Username string
	User_picture string
}

func GetRoomComment(room_id string) []*RoomComments  {
	o:=orm.NewOrm()
	o.Using("default")
	var comments []*RoomComments
	_,err :=o.Raw("SELECT a.*,b.username,c.user_picture from room_comments as a left join user as b  on a.user_id=b.id left join user_info as c on b.id=c.user_id where a.room_id = ?",room_id).QueryRows(&comments)
    if err !=nil {
    	fmt.Println("查询失败",err)
	}
	return comments
}