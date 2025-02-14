package globals

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	UNICRM_REDIS      redis.UniversalClient
	UNICRM_REDIS_LIST map[string]redis.UniversalClient
)

// GetRedis 获取Redis配置
func GetRedis(name string) redis.UniversalClient {
	redis, ok := UNICRM_REDIS_LIST[name]
	if redis == nil || !ok {
		panic(fmt.Sprintf("未找到Redis配置: %s", name))
	}
	return redis
}
