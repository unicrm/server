package logging

import (
	"time"

	"go.uber.org/zap/zapcore"
)

var (
	LoggerConfig Logger
)

// Logger 日志配置结构体
type Logger struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Director      string `mapstructure:"director" json:"director" yaml:"director"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"` // 输出日志的格式，可选的有json和console
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"`
	MaxAge        int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
}

// 需要输出的日志级别，Logger.Level 以上的日志级别才会输出，例如：info,warn,error,fatal
func (logger *Logger) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 7)
	level, err := zapcore.ParseLevel(logger.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	for ; level <= zapcore.FatalLevel; level++ {
		levels = append(levels, level)
	}
	return levels
}

// ConsoleEncoder 控制台输出格式配置
func (logger *Logger) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		MessageKey:    "message",
		TimeKey:       "time",
		LevelKey:      "level",
		CallerKey:     "caller",
		NameKey:       "name",
		StacktraceKey: logger.StacktraceKey,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    logger.LevelEncoder(),
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// 根据配置文件中的格式选择输出日志的格式
	if logger.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

// 日志级别配置
func (logger *Logger) LevelEncoder() zapcore.LevelEncoder {
	return zapcore.LowercaseLevelEncoder
}
