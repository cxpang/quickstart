package models

import (
	"github.com/astaxie/beego/orm"
)

//数据room字段映射
type Room struct {
	Room_id int  `orm:"pk;column(room_id);"`
	User_id int
	Room_name string
	Room_type int
	Room_price int
	Room_content string
	Room_detail string
	Room_images string
	Room_city int
	Room_area int
	Room_address string
	Add_time string
	Update_time string
	Answer_users int
	Focus_count int
	Comment_count int
	Is_essence int
	Is_comment int
	Is_over int
	Point_lng float64
	Point_lat float64
	Is_check int
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Room))
}

func GetAllRooms() []*Room {
	o:=orm.NewOrm()
	o.Using("default")
	var rooms []*Room
	q:= o.QueryTable("room")
	q.All(&rooms)
	return rooms
}
