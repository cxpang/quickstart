package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"

	"encoding/json"
	"fmt"
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
var users map[string]string


func (this *WebsocketController) Join() {
	userid := this.GetString("UserId") //获取发起私聊的用户id
	touserid:=this.GetString("ToUserId")
	if len(userid) == 0 {
		this.Redirect("/", 302)
		return
	}
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	WebsocketJoin(userid)
	CheckUserisLogin(touserid,ws)

	//客户端离开事件
	defer Leave(userid)


	//处理聊天消息

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(p))
	}





}

//利用写协程发送消息


func  WebsocketJoin(userid string){
	 users = make(map[string]string)
	 _,ok:=users[userid]
	 if !ok {
	 	users[userid] = userid
	 }
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
	delete(users, userid)
}