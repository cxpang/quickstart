package controllers
import (
	"github.com/astaxie/beego"

	"quickstart/models"
   "fmt"
)
type LoginController struct {
	beego.Controller
}
func (this *LoginController) Get(){
	this.Layout = "layout/layout.html"
	this.TplName = "site/login.html"
}
func (this *LoginController) Post(){
	email:=this.GetString("email")
	username:=this.GetString("username")
    datas,isbool:=models.CheckUserLogin(username,email)
    fmt.Println(datas)
    if isbool==false{
		this.Redirect("/login",302)
	}
	this.Redirect("/",301)

}