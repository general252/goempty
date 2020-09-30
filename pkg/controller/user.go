package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterUser(router *gin.RouterGroup) {
	user := userController{}

	router.POST("/user/login", user.login)
	router.POST("/user/logout", user.logout)
}

type userController struct {
}

func (*userController) login(c *gin.Context) {

}

func (*userController) logout(c *gin.Context) {

}
