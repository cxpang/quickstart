package models
import (
	"github.com/astaxie/beego/orm"

)

type User struct {
	Id int  `orm:"pk;column(id);"`
	Username string
	Email string
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
func CheckUserLogin(username,email string) (User,bool)  {

	o:=orm.NewOrm()
	o.Using("default")
	var user User
	err :=o.QueryTable("user").Filter("username",username).Filter("email",email).One(&user)
	if err !=nil{
		return user,false
	}
	return user,true
}

