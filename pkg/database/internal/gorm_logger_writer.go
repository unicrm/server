package internal

import (
	"fmt"

	"github.com/unicrm/server/pkg/database/models"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config models.GeneralDB
}

func NewWriter(config models.GeneralDB) *Writer {
	return &Writer{config: config}
}

// Printf 格式化打印日志
func (w *Writer) Printf(message string, data ...any) {
	// zap日志输出
	switch w.config.LogLevel() {
	case logger.Silent:
		zap.L().Debug(fmt.Sprintf(message, data...))
	case logger.Error:
		zap.L().Error(fmt.Sprintf(message, data...))
	case logger.Warn:
		zap.L().Warn(fmt.Sprintf(message, data...))
	case logger.Info:
		zap.L().Info(fmt.Sprintf(message, data...))
	default:
		zap.L().Info(fmt.Sprintf(message, data...))
	}
}
