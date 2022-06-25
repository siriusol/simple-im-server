package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

func StartWebSocket() {
	http.HandleFunc("/acc", wsPage)
	http.ListenAndServe(":8089", nil)
}

var clientManager ClientManager

type login struct {
}

func wsPage(w http.ResponseWriter, req *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referfer:", r.Header["Referer"])
			return true
		},
	}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	fmt.Println("websocket 建立连接", conn.RemoteAddr().String())

	currentTime := time.Now().Unix()
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}

type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户，appId+uuid
	UserLock    sync.RWMutex
	Register    chan *Client // 连接处理
	Login       chan *login  // 用户登录处理
	Unregister  chan *Client // 断开连接处理程序
	Broadcast   chan []byte  // 广播 向全部成员发送数据
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}
}

func ProcessData(c *Client, message []byte) {}
