package logging

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerInit(config Logger) *zap.Logger {

	// 初始化配置
	LoggerConfig = Logger(config)

	// 判断是否有Director文件夹
	director := LoggerConfig.Director
	if ok, _ := PathExists(director); !ok {
		if err := CreateDir(director); err != nil {
			panic(fmt.Errorf("初始化日志异常: %v", err))
		}
	}

	// 初始化日志配置文件
	levels := LoggerConfig.Levels()
	cores := make([]zapcore.Core, 0, len(levels))
	for _, level := range levels {
		core := NewZapCore(level)
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))

	// 添加堆栈信息，只在error级别时添加堆栈信息
	// logger = logger.WithOptions(zap.AddStacktrace(zapcore.ErrorLevel))

	// 添加调用者信息
	if LoggerConfig.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
