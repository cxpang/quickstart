package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Commentres struct {
	Id int `orm:"pk;column(id);"`
	Comments_id int
	User_id int
	At_user_id int
	Message string
	Created_at int
	Username string
	At_username string
	User_picture string
	At_user_picture string
}
func GetRoomCommentRes(comment_id int) []*Commentres{

	var commentres  []*Commentres
	o:=orm.NewOrm()
	o.Using("default")
	_,err := o.Raw(" select * from room_comments_res where comments_id = ? ",comment_id).QueryRows(&commentres)
	if err !=nil{
		fmt.Println(err)
	}
	return commentres
}
