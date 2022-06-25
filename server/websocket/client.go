package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

type Client struct {
	Addr   string
	Socket *websocket.Conn
	Send   chan []byte
}

func NewClient(addr string, socket *websocket.Conn, firstTime int64) *Client {
	return &Client{}
}

func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		clientManager.Unregister <- c
		c.Socket.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误，关闭连接
				fmt.Println("Client发送数据，关闭连接", c.Addr, "ok", ok)
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("read stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		fmt.Println("读取客户端数据 关闭send", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("读取客户端数据错误", c.Addr, err)
			return
		}

		// 处理程序
		fmt.Println("读取客户端数据处理", string(message))
		ProcessData(c, message)
	}
}
