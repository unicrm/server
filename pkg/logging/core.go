package logging

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

func NewZapCore(level zapcore.Level) zapcore.Core {
	entity := &ZapCore{level: level}
	encoder := LoggerConfig.Encoder()
	// 日志写入器
	syncer := entity.WriteSyncer()
	// 日志级别过滤器
	enabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == entity.level
	})
	// 初始化zapcore.Core
	entity.Core = zapcore.NewCore(encoder, syncer, enabler)
	return entity.Core
}

// WriteSyncer
// 自定义zapcore.WriteSyncer，这里用的是lumberjack.Logger
// lumberjack.Logger实现了zapcore.WriteSyncer接口
func (zc *ZapCore) WriteSyncer() zapcore.WriteSyncer {

	// 日志文件名
	filename := "server." + zc.level.String() + ".log"

	// 日志文件所在的目录
	values := []string{LoggerConfig.Director, filename}

	// 初始化lumberjack.Logger
	hook := &lumberjack.Logger{
		Filename:   filepath.Join(values...), // 日志文件的位置
		MaxSize:    10,                       // 日志文件在写到这么大的时候将执行滚动替换
		MaxBackups: 3,                        // 保留旧文件的最大个数
		MaxAge:     LoggerConfig.MaxAge,      // 保留旧文件的最大天数
	}

	// 是否同时输出到控制台
	if LoggerConfig.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook), zapcore.AddSync(os.Stdout))
	}
	return zapcore.AddSync(hook)
}
