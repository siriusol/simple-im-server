package router

import "github.com/siriusol/simple-im-server/server/websocket"

func WebsocketInit() {
	websocket.Register("login", websocket.Login)
	websocket.Register("heartbeat", websocket.Heartbeat)
	websocket.Register("ping", websocket.Ping)
}
