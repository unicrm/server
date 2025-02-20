package main

import (
	"github.com/unicrm/server/internal/core"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/initialize"
	"github.com/unicrm/server/pkg/database"
	"github.com/unicrm/server/pkg/logging"
	"github.com/unicrm/server/pkg/redis"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

func main() {
	// 加载配置文件
	initialize.LoadConfig()
	// 初始化日志
	globals.UNICRM_LOGGER = logging.LoggerInit(globals.UNICRM_CONFIG.Logger)
	zap.ReplaceGlobals(globals.UNICRM_LOGGER) // 替换全局的日志对象
	// 初始化认证服务
	initialize.OtherInit()
	// 注册定时任务
	initialize.RegisterTimer()
	// 初始化数据库
	globals.UNICRM_DB = database.DatabaseInit(globals.UNICRM_CONFIG.GeneralDB)
	defer database.DatabaseClose(globals.UNICRM_DB)
	// 注册数据库表结构
	database.RegisterTables(globals.UNICRM_DB, globals.UNICRM_TABLES...)
	// 初始化Redis
	globals.UNICRM_REDIS_LIST = redis.InitRedisList(globals.UNICRM_CONFIG.RedisList)
	globals.UNICRM_REDIS = globals.GetRedis(globals.UNICRM_CONFIG.System.RedisName)
	// 启动服务
	core.RunServer()
}
