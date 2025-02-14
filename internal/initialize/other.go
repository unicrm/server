package initialize

import (
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/pkg/auth/tools"
	"go.uber.org/zap"
)

func OtherInit() {
	if _, err := tools.ParseDuration(globals.UNICRM_CONFIG.JWT.ExpiresTime); err != nil {
		zap.L().Panic("解析JWT过期时间失败", zap.Error(err))
	}
	if _, err := tools.ParseDuration(globals.UNICRM_CONFIG.JWT.BufferTime); err != nil {
		zap.L().Panic("解析JWT缓冲时间失败", zap.Error(err))
	}
}

func init() {}
