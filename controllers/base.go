package controllers
import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	username string
	userpicture string
	IsLogin  int
}
func (this BaseController) Prepare(){
	username:=this.GetSession("user_name")
	userpicture:=this.GetSession("user_picture")
	if username==nil && userpicture==nil {
		this.Data["IsLogin"]=1
		this.Data["User_name"]="未登录"
		this.Data["User_picture"]="/roomshare/uploads/null.jpg"
	}else {
		this.Data["IsLogin"]=2
		this.Data["User_name"]=username
		this.Data["User_picture"]=userpicture
	}
}
