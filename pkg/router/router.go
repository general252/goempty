package router

import (
	"github.com/general252/goempty/pkg/controller"
	"github.com/general252/goempty/pkg/middleware"
	"github.com/general252/goempty/pkg/util"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	// 中间件
	router.Use(middleware.Cors())

	v1 := router.Group("/empty/v1")
	// websocket
	v1.GET("/ws/:TOKEN", util.GinWebsocketHandler(controller.WebsocketConnHandle))

	// 用户API
	userApi := v1.Group("/user")
	{
		user := controller.NewUserController()

		userApi.POST("/login", user.Login)
		userApi.POST("/logout", user.Logout)

		userApi.POST("/", user.AddUser)          // add
		userApi.DELETE("/:UserId", user.DelUser) // delete
		userApi.PUT("/:UserId", user.UpdateUser) // update
		userApi.GET("/:UserId", user.GetUser)    // get
		userApi.POST("/query", user.Query)       // query
	}

	return router
}
