package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"time"
)

func WebsocketConnHandle(ws *websocket.Conn) {
	defer func() {
		_ = ws.Close()
	}()

	var msg interface{}
	for {
		err := websocket.JSON.Receive(ws, msg)
		if err != nil {
			break
		}

		err = websocket.JSON.Send(ws, gin.H{
			"date": time.Now(),
		})
		if err != nil {
			break
		}
	}
}
