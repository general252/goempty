package pkg

import (
	"github.com/general252/goempty/pkg/db"
	"github.com/general252/goempty/pkg/parseSwag"
	"github.com/general252/goempty/pkg/router"
	"github.com/general252/goempty/pkg/version"
	"github.com/gin-gonic/gin"
)

func VMain() {
	// 整理swag api
	parseSwag.ParseSwag()

	// 显示版本信息
	version.ShowVersionInfo()

	// 初始化数据库
	_ = db.InitDataBase()

	// 初始化redis

	var r = gin.Default()
	router.InitRouter(r)

	// 启动服务

	// 等待事件

	// 关闭服务
}
