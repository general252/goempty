package util

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

// GinWebsocketHandler websocket.Handler 转 gin HandlerFunc
func GinWebsocketHandler(wsConnHandle websocket.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.IsWebsocket() {
			wsConnHandle.ServeHTTP(c.Writer, c.Request)
		} else {
			_, _ = c.Writer.WriteString("===not websocket request===")
		}
	}
}
