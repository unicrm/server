package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisList map[string]redis.UniversalClient

func initRedisClient(redisConfig Redis) (redis.UniversalClient, error) {

	var client redis.UniversalClient
	if redisConfig.UseCluster {
		// 集群模式
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisConfig.ClusterAddrs,
			Password: redisConfig.Password,
		})
	} else {
		// 单机模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
	}

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("REDIS连接失败", zap.String("pong", pong), zap.Error(err))
		return nil, err
	}
	zap.L().Info("REDIS连接成功", zap.String("pong", pong), zap.String("name", redisConfig.Name))
	return client, nil
}

func InitRedisList(redisConfigList []Redis) RedisList {
	redisMap := make(map[string]redis.UniversalClient)
	for _, redisConfig := range redisConfigList {
		redisClient, err := initRedisClient(redisConfig)
		if err != nil {
			zap.L().Panic("REDIS初始化失败", zap.Error(err))
		}
		redisClient.AddHook(RedisHook{})
		redisMap[redisConfig.Name] = redisClient
	}
	return redisMap
}
