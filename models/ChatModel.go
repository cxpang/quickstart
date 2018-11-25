package models
import (
	"github.com/astaxie/beego/orm"

	"time"
)

type Chatpoint struct {
	Id int  `orm:"pk;column(id);"`
	Fromid int
	Toid  int
	CreatedAt int64
	UpdatedAt int64
	Message  string
 	IsRead string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Chatpoint))
}
func CheckUserChatIsExist(FromUserId int,ToUserId int) bool{
	o:=orm.NewOrm()
	o.Using("default")
    var chatUser Chatpoint
	err :=o.QueryTable("chatpoint").Filter("fromid",FromUserId).Filter("toid",ToUserId).One(&chatUser)
	if err!=nil{
		return false
	}
	return true

}
func SaveChatInfo(FromUserId int,ToUserId int){
	o:=orm.NewOrm()
	o.Using("default")
    chatpoint :=new (Chatpoint)
    chatpoint.Fromid = FromUserId
    chatpoint.Toid = ToUserId
    chatpoint.CreatedAt = time.Now().Unix()
    chatpoint.UpdatedAt = time.Now().Unix()
    chatpoint.IsRead = "0";
	o.Insert(chatpoint)
}
//获取当前用户的所有聊天人信息 来自聊天
func GetALLChatFromInfo(currentId int)([]*Chatpoint,bool){
	o:=orm.NewOrm()
	o.Using("default")
	var chatUsers []*Chatpoint
	_,err := o.QueryTable("chatpoint").Filter("fromid",currentId).All(&chatUsers)
	if err!=nil{
		return chatUsers,true
	}
	return chatUsers,false
}
//获取当前用户的所有聊天人信息 发起聊天
func GetALLChatToInfo(currentId int)([]*Chatpoint,bool){
	o:=orm.NewOrm()
	o.Using("default")
	var chatUsers []*Chatpoint
	_,err := o.QueryTable("chatpoint").Filter("toid",currentId).All(&chatUsers)
	if err!=nil{
		return chatUsers,true
	}
	return chatUsers,false
}