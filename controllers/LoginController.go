package controllers
import (
	"github.com/astaxie/beego"

	"quickstart/models"
    //"fmt"
)
type LoginController struct {
	beego.Controller
}
func (this *LoginController) Get(){

	this.TplName = "site/login.html"
}
func (this *LoginController) Post(){
	email:=this.GetString("email")
	username:=this.GetString("username")
    datas,isbool:=models.CheckUserLogin(username,email)
    if isbool==false{
		this.Redirect("/login",302)
	}
	this.SetSession("user_name",datas.Username)
	this.SetSession("user_picture",datas.User_picture)
	//fmt.Println(this.CruSession)
	this.Redirect("/",302)
}