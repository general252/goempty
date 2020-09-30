package router

import (
	"github.com/general252/goempty/pkg/controller"
	"github.com/general252/goempty/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	middleware.UseMiddleWare(router)

	v1 := router.Group("/v1")
	controller.RegisterUser(v1)
	controller.RegisterPlugin(v1)

	controller.RegisterWebSocket(v1)
}
