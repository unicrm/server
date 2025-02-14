package redis

import (
	"context"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var ReadisList []Redis

func TestMain(m *testing.M) {
	// 初始化日志
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	// 初始化配置文件
	v := viper.New()
	v.SetConfigFile("../../config.debug.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logger.Error("读取配置文件失败", zap.Error(err))
	}
	v.WatchConfig()
	if err := v.UnmarshalKey("redis-list", &ReadisList); err != nil {
		logger.Error("解析配置文件失败", zap.Error(err))
	}
	m.Run()
}

func TestRedisList(t *testing.T) {

	redisList := InitRedisList(ReadisList)

	defer redisList["default"].Close()
	redisList["default"].Set(context.Background(), "test", "test", 0)
	test1 := redisList["default"].Get(context.Background(), "test")
	assert.Equal(t, "test", test1.Val())

	redis := redisList[DEFAULT_REDIS_NAME]
	defer redis.Close()
	redis.Set(context.Background(), "test", "test", 0)
	test2 := redis.Get(context.Background(), "test")
	assert.Equal(t, "test", test2.Val())

	defer redisList["cache"].Close()
	redisList["cache"].Set(context.Background(), "test", "test", 0)
	test3 := redisList["cache"].Get(context.Background(), "test")
	assert.Equal(t, "test", test3.Val())
}
