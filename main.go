package main

import (
	"github.com/siriusol/simple-im-server/server/websocket"
)

func main() {
	go websocket.StartWebSocket()
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	//currentTime := uint64(time.Now().Unix())
	//for client := range clientManager.Clients {
	//	if client.IsHeartbeatTimeout(currentTime) {
	//		fmt.Println("心跳时间超时，关闭连接", client.Addr, client.UserId, clientLoginTime, client.HeartbeatTime)
	//		client.Socket.Close()
	//	}
	//}
}
