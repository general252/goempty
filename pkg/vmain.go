package pkg

import (
	"context"
	"fmt"
	"github.com/general252/goempty/pkg/config"
	"github.com/general252/goempty/pkg/model"
	"github.com/general252/goempty/pkg/router"
	"github.com/general252/goempty/pkg/version"
	"github.com/general252/gout/ulog"
	"github.com/general252/gout/usafe"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func VMain() {
	// 显示版本信息
	version.ShowVersionInfo()

	// 初始化数据库
	if err := model.InitDataBase(config.JsonConfig.DB.ConnectionString()); err != nil {
		ulog.Error("init database %v", err)
		return
	}

	// 监听端口
	var serverPort = config.JsonConfig.HttpPort
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", serverPort))
	if err != nil {
		ulog.Error("listen %v fail %v", serverPort, err)
		return
	}

	var httpServer = &http.Server{
		Handler: router.InitRouter(),
	}

	// 启动服务
	var pool = usafe.NewPool(context.TODO())
	pool.GoCtx(func(ctx context.Context) {
		if err := httpServer.Serve(listener); err != nil {
			ulog.Error("serve fail %v", err)
		}
	})

	// 等待事件
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-quitChan

	_ = httpServer.Shutdown(context.TODO())

	// 关闭服务
	pool.Stop()
}
