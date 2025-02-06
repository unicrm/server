package core

import (
	"fmt"
	"net/http"

	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/initialize"
	"github.com/unicrm/server/internal/services/system"

	"go.uber.org/zap"
)

func RunServer() {

	// 从数据库加载jwt黑名单
	system.LoadJwtBlackList()

	// 初始化路由
	Router := initialize.Routers()
	address := fmt.Sprintf(":%s", globals.UNICRM_CONFIG.System.Addr)
	srv := initServer(address, Router)
	globals.UNICRM_LOGGER.Info("服务启动成功，监听地址：", zap.String("address", address))

	fmt.Printf("\n")
	fmt.Printf("使用说明 \n")
	fmt.Printf("初始化文档: swag init -g cmd/server/main.go \n")
	fmt.Printf("文档访问地址: \033[32m%s/swagger/index.html \033[0m\n", address)
	fmt.Printf("\n")

	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			globals.UNICRM_LOGGER.Error("服务启动失败", zap.Error(err))
		}
	}()
	// globals.UNICRM_LOGGER.Error("服务启动失败", zap.Error(srv.ListenAndServe()))

	// 优雅关闭
	closeServerLocked(srv)

}
