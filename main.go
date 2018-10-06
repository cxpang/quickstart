package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)
//自定义模板函数实现分割字符串
func SubstrForTemplate(str string) (res string)  {
	 resu1:=strings.Split(str, ",")
	 return resu1[0]
}

func main() {
	beego.AddFuncMap("SubstrForTemplate", SubstrForTemplate)
	beego.Run()
}
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:200517@tcp(116.62.232.139:9090)/roomshare?charset=utf8")
}


