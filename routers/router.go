package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/roomall", &controllers.MainController{})
	beego.Router("/roomdetail/:roomid:int", &controllers.RoomdetailController{})
}
