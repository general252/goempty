package controller

import "github.com/gin-gonic/gin"

func RegisterPlugin(router *gin.RouterGroup) {
	plugin := pluginController{}

	router.POST("/plugin/add", plugin.addPlugin)
	router.POST("/plugin/del", plugin.delPlugin)
}

type pluginController struct {
}

func (*pluginController) addPlugin(c *gin.Context) {

}

// @summary 03-02 删除插件版本
// @Description 删除插件版本
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param KeyId path string true "插件版本KeyId(唯一标志)"
// @Success 200 {object}  JsonResult	"ok"
// @Router /plugin/del/{KeyId} [delete]
func (*pluginController) delPlugin(c *gin.Context) {

}
