package middleware

import (
	"github.com/gin-gonic/gin"
)

func UseMiddleWare(router *gin.Engine) {
	router.Use(cors())
}
