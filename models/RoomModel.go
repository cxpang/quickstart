package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
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
	User_picture string
	City string
	Area string
}
type RoomDetail struct {
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
	Add_time int64
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
	User_picture string
	City string
	Area string
	Username string
	Email string
	User_expe int
	User_credit int
	User_university int
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Room))
}

func GetAllRooms() []*Room {
	o:=orm.NewOrm()
	o.Using("default")
	var rooms []*Room
	//q:= o.QueryTable("room").Filter("is_check",1)
	//q.All(&rooms)
	_,err :=o.Raw("SELECT a.*,b.user_picture,c.city,d.area FROM room AS a LEFT JOIN user_info AS b ON a.user_id = b.user_id LEFT JOIN city AS c ON a.room_city = c.city_id LEFT JOIN area AS d ON a.room_area = d.area_id WHERE a.is_check = ?",1).QueryRows(&rooms)

	if err ==nil{
       return rooms
	}
	return rooms
}
func GetRoomDetail(roomid string) RoomDetail{
	o:=orm.NewOrm()
	o.Using("default")
	var roomDetail RoomDetail
	err := o.Raw("SELECT a.*,b.user_picture,b.user_university,b.user_expe,b.user_credit,c.city,d.area,e.username FROM room AS a LEFT JOIN user_info AS b ON a.user_id = b.user_id LEFT JOIN user as e ON a.user_id=e.id LEFT JOIN city AS c ON a.room_city = c.city_id LEFT JOIN area AS d ON a.room_area = d.area_id WHERE a.room_id = ?",roomid).QueryRow(&roomDetail)
	if err !=nil{
		fmt.Println(err)
	}
	return roomDetail
}
