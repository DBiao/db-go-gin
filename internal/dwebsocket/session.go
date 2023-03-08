package dwebsocket

import (
	"db-go-gin/internal/utils"
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Out  chan []byte
}

func (c *Client) Read(ws *websocket.Conn) {
	defer utils.PrintPanic()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (c *Client) Write(ws *websocket.Conn) {
	defer utils.PrintPanic()

	for {
		select {
		case out, ok := <-c.Out:
			if !ok {
				break
			}
			err := ws.WriteMessage(websocket.TextMessage, out)
			if err != nil {
				break
			}
		}
	}
}

// Close 关闭监听
func (c *Client) Close(ws *websocket.Conn) {
	ws.SetCloseHandler(func(code int, text string) error {
		Delete(c)
		ws.Close()
		fmt.Println("断开成功", code, text)
		return nil
	})
}
