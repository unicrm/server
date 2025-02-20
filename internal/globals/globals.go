package globals

import (
	"github.com/patrickmn/go-cache"
	"github.com/unicrm/server/internal/config"
	"github.com/unicrm/server/internal/utils/timer"
	"github.com/unicrm/server/pkg/auth"
	"go.uber.org/zap"
)

var (
	UNICRM_CONFIG      *config.UnicrmConfig
	UNICRM_LOGGER      *zap.Logger
	UNICRM_TIMER       timer.Timer = timer.NewTimerTask()
	UNICRM_AUTH        *auth.AuthExtend
	UNICRM_BLACK_CACHE *cache.Cache
)
