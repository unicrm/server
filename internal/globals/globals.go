package globals

import (
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"github.com/unicrm/server/internal/config"
	"github.com/unicrm/server/internal/utils/timer"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// 全局配置
	UNICRM_CONFIG     *config.UnicrmConfig
	UNICRM_LOGGER     *zap.Logger
	UNICRM_DB         *gorm.DB
	UNICRM_TIMER      timer.Timer = timer.NewTimerTask()
	UNICRM_REDIS_LIST map[string]redis.UniversalClient

	// 黑名单缓存
	UNICRM_BLACK_CACHE *cache.Cache
)
