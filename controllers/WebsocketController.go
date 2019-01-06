package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"

	"encoding/json"
	"fmt"
	"log"
)

type WebsocketController struct {
	beego.Controller
}
type Message struct {
     Status int
     Data string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//定义用户结构体
type Usersocket struct {
	UserId  string
	ws  *websocket.Conn
}
//定义所有在线用户数组
var users = make(map[string]Usersocket)

type UserMessage struct {
	Content string
	Fromuser  string
	Touser string
}
var publish = make(chan UserMessage, 10)


func (this *WebsocketController) Join() {
	userid := this.GetString("UserId") //获取发起私聊的用户id
	touserid:=this.GetString("ToUserId")
	if len(userid) == 0 {
		this.Redirect("/", 302)
		return
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request,nil)
	if err != nil {
		log.Fatal(err)
	}
	WebsocketJoin(ws,userid)



    defer Leave(userid)
    defer ws.Close()
	//处理聊天消息

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {

			return
		}
		//string(p)为接收到的消息
		var usermess UserMessage
		usermess.Fromuser = userid
		usermess.Touser = touserid
		usermess.Content = string(p)
		publish <- usermess
		go handlechat()
	}





}

//利用写协程发送消息

func handlechat(){
	chatmess:=<-publish
	tousers := chatmess.Touser
	fromuser :=chatmess.Fromuser
	content:=chatmess.Content
	fromuserWs :=users[fromuser].ws
	onlineuser ,ok:= users[tousers]
	if !ok {
		var  mess Message
		mess.Status=3
		mess.Data = "该用户不存在"
		response,_ := json.Marshal(mess)
		fromuserWs.WriteMessage(1,response)
	} else {
		//用户在线处理发送
		tousersws := onlineuser.ws
		var mess Message
		mess.Status = 1
		mess.Data = content
		response, _ := json.Marshal(mess)
		tousersws.WriteMessage(1, response)
	}
}



func  WebsocketJoin(ws *websocket.Conn,userid string){

	 _,ok:=users[userid]
	 if !ok {
	 	users[userid]=Usersocket{userid,ws}
	 }
	 fmt.Println("用户加入聊天，用户id",users)
	 return
}
/**
 *查看users数组中是否存在要私聊的用户
  不存在发送给客户端
*/
func CheckUserisLogin(touserid string,ws  *websocket.Conn){
	_,ok:=users[touserid]
	if !ok {
        var  mess Message
        mess.Status=3
        mess.Data = "用户不存在"
        response,_ := json.Marshal(mess)
        ws.WriteMessage(1,response)

	}
}
//用户断开
func Leave(userid string){
	fmt.Println("用户离开聊天，用户id",users)
	delete(users, userid)
}