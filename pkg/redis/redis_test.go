package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unicrm/server/pkg/redis/internal"
)

func TestRedisList(t *testing.T) {
	redisConfigList := []Redis{
		{
			Name:         "default",
			Addr:         "106.12.59.2:6379",
			Password:     "123456",
			DB:           0,
			UseCluster:   false,
			ClusterAddrs: []string{},
		},
		{
			Name:         "cache",
			Addr:         "106.12.59.2:6379",
			Password:     "123456",
			DB:           0,
			UseCluster:   false,
			ClusterAddrs: []string{},
		},
	}

	redisList := InitRedisList(redisConfigList)

	defer redisList["default"].Close()
	redisList["default"].Set(context.Background(), "test", "test", 0)
	test1 := redisList["default"].Get(context.Background(), "test")
	assert.Equal(t, "test", test1.Val())

	redis := redisList[internal.DEFAULT_REDIS_NAME]
	defer redis.Close()
	redis.Set(context.Background(), "test", "test", 0)
	test2 := redis.Get(context.Background(), "test")
	assert.Equal(t, "test", test2.Val())

	defer redisList["cache"].Close()
	redisList["cache"].Set(context.Background(), "test", "test", 0)
	test3 := redisList["cache"].Get(context.Background(), "test")
	assert.Equal(t, "test", test3.Val())
}
