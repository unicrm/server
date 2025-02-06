package core

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

func initServer(address string, router *gin.Engine) server {
	// 初始化服务端
	srv := endless.NewServer(address, router)
	srv.ReadHeaderTimeout = 10 * time.Minute
	srv.WriteTimeout = 10 * time.Minute
	srv.MaxHeaderBytes = 1 << 20
	return srv
}

func closeServerLocked(srv server) {
	// 关闭服务端
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		globals.UNICRM_LOGGER.Error("服务关闭失败", zap.Error(err))
	}
}
