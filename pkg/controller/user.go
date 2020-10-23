package controller

import (
	"fmt"
	"github.com/general252/goempty/pkg/model/dao"
	"github.com/general252/gout/ulog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserController() *userControllerWrap {
	return &userControllerWrap{}
}

type userControllerWrap struct {
	userController
}

func (tis *userControllerWrap) Login(c *gin.Context) {
	ulog.Info("============== before ")
	c.Set("see_you", "1314")

	tis.userController.Login(c)

	data, exists := c.Get("see_you")
	ulog.Info("============== after %v %v", data, exists)
}

type userController struct {
}

type JsonLoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (*userController) Login(c *gin.Context) {
	var result = &JsonResult{
		Code: -1,
		Msg:  "fail",
	}

	for {
		// 检查参数
		var param JsonLoginReq
		if err := c.ShouldBindJSON(&param); err != nil {
			result.Code = StatusCodeFail
			result.Msg = err.Error()
			break
		}

		// 查询数据库
		var helpUser = dao.NewUserDao()
		objUser, err := helpUser.Get(param.UserName)
		if err != nil {
			// 数据库错误
			result.Code = StatusCodeErrorDataBase
			result.Msg = err.Error()
			break
		}
		// 没有查到
		if objUser == nil {
			result.Code = StatusCodeNotFound
			result.Msg = "not found user"
			break
		}

		// 判断密码
		if objUser.Password != param.Password {
			result.Code = StatusCodeErrorPassword
			result.Msg = "error password"
			break
		}

		// 成功
		result.Code = StatusCodeOK
		result.Msg = fmt.Sprintf("%v login success", objUser.UserName)

		break
	}

	data, exists := c.Get("see_you")
	ulog.Info("-------------- %v %v", data, exists)

	c.JSON(http.StatusOK, result)
}

func (*userController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (*userController) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (*userController) DelUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (*userController) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (*userController) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (*userController) Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
