package main

import (
	//"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	//fmt.Println("hello world")

	server, err := socketio.NewServer(nil) //新建一个server
	if err != nil {
		log.Fatal(err)
	}

	//server.On("connection", func(so socketio.Socket){
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		log.Println("connected:", so.ID())

		// so.join("chat")
		// so.On("chat message", func(msg string){
		// 	log.Println("Message received from Client:" + msg)
		// 	so.BroadcastTo("chat","chat message", msg)
		// })
		return nil
	})

	server.OnEvent("/chat", "msg", func(so socketio.Conn, msg string) string {
		so.SetContext(msg)
		return "recv " + msg
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs) //主页访问让fs必要时处理

	http.Handle("/socket.io/", server) //前端html页面会访问这个url
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
