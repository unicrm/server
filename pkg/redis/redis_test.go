package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	redisListConfig []Redis
	redisList       RedisList
)

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
	if err := v.UnmarshalKey("redis-list", &redisListConfig); err != nil {
		logger.Error("解析配置文件失败", zap.Error(err))
	}
	m.Run()
}

func TestRedisList(t *testing.T) {

	redisList = InitRedisList(redisListConfig)
	fmt.Println(redisList)

	redisList[DEFAULT_REDIS_NAME].Set(context.Background(), "test", "test", 0)
	test := redisList[DEFAULT_REDIS_NAME].Get(context.Background(), "test")
	assert.Equal(t, "test", test.Val())

	redisList["cache"].Set(context.Background(), "cache", "cache", 0)
	cache := redisList["cache"].Get(context.Background(), "cache")
	assert.Equal(t, "cache", cache.Val())
}

func TestRedisNil(t *testing.T) {
	ss, err := redisList[DEFAULT_REDIS_NAME].Get(context.Background(), "ss").Result()
	fmt.Println(ss, err)

	aa := redisList[DEFAULT_REDIS_NAME].Get(context.Background(), "aa").Val()
	fmt.Println(aa)
}

func TestRedisClose(t *testing.T) {
	defer redisList["default"].Close()
	defer redisList["cache"].Close()
}
