package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"Demo-in-Golang/demo/websocket/impl"
)

/**
 * @Author  jackie.lqj
 * @Date  2022/5/18 19:36
 * @Description 封装websocket
 */

var (
	upgrade = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// websocket 长连接
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	// header中添加Upgrade:websocket
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		log.Println("upgrade error:", err)
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		log.Println("InitConnection error:", err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			if err := conn.WriteMessage([]byte("heartbeat")); err != nil {
				log.Println("write heartbeat error:", err)
				return
			}
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			log.Printf("read message error: %s", err.Error())
			return
		}
		log.Printf("got message: %s", data)
		if err = conn.WriteMessage(data); err != nil {
			log.Printf("write message error: %s", err.Error())
			return
		}
	}
}

func main() {
	// http标准库
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe("0.0.0.0:5555", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
