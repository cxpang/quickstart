package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"quickstart/models"
	"time"
)
//自定义模板函数实现分割字符串
func SubstrForTemplate(str string) (res string)  {
	 resu1:=strings.Split(str, ",")
	 return resu1[0]
}
func GetUserName(userid int)(res string){
	Users,_:=models.GetUserNameByUserId(userid)
	return Users.Username
}
func GetUserPicture(userid int)(res string){
	Users,_:=models.GetUserNameByUserId(userid)
	return Users.User_picture
}
//格式化int64时间为H-m-d H:i:s
func FormateTimeFromInt64(num int64) string{
	startDate := time.Unix(num, 0).Format("2006-01-02 15:04:05")
	return  startDate
}
func main() {
	beego.AddFuncMap("SubstrForTemplate", SubstrForTemplate)
	beego.AddFuncMap("GetUserName", GetUserName)
	beego.AddFuncMap("GetUserPicture", GetUserPicture)
	beego.AddFuncMap("FormateTimeFromInt64", FormateTimeFromInt64)
	beego.Run()
}
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:200517@tcp(116.62.232.139:9090)/roomshare?charset=utf8")
}


