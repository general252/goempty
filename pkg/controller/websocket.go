package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"time"
)

func RegisterWebSocket(router *gin.RouterGroup) {
	router.GET("/ws/:TOKEN", ginWebsocketHandler(websocketConnHandle))
}

func websocketConnHandle(ws *websocket.Conn) {
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

// websocket.Handler 转 gin HandlerFunc
func ginWebsocketHandler(wsConnHandle websocket.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.IsWebsocket() {
			wsConnHandle.ServeHTTP(c.Writer, c.Request)
		} else {
			_, _ = c.Writer.WriteString("===not websocket request===")
		}
	}
}
