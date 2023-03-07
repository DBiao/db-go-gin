package router

import (
	"db-go-gin/internal/dwebsocket"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandle(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	client := &dwebsocket.Client{
		Conn: ws,
		Out:  make(chan []byte, 10),
	}

	fmt.Println("连接成功")

	dwebsocket.Add(client)

	ws.SetPingHandler(func(appData string) error {
		fmt.Println("ping")
		ws.WriteControl(websocket.PongMessage, []byte(appData), time.Time{})
		return nil
	})

	// 开启关闭连接监听
	client.Close(ws)

	go client.Read(ws)
	go client.Write(ws)
}
