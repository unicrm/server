package globals

import (
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/unicrm/server/internal/models/system"
	"github.com/unicrm/server/pkg/database/model"
)

const (
	InitOrderSystem = iota
)

var (
	RegisterTables []interface{}
)

var (
	ErrMissingDBContext = errors.New("未找到数据库上下文")
)

type InitContentKey string

// GetGeneralDB 获取通用数据库配置
func GetGeneralDB() model.GeneralDB {
	var generalDB model.GeneralDB
	switch UNICRM_CONFIG.System.DBType {
	case "mysql":
		generalDB = UNICRM_CONFIG.Mysql.GeneralDB
		generalDB.DBType = "mysql"
		return generalDB
	default:
		generalDB = UNICRM_CONFIG.Mysql.GeneralDB
		generalDB.DBType = "mysql"
		return generalDB
	}
}

// GetRedis 获取Redis配置
func GetRedis(name string) redis.UniversalClient {
	redis, ok := UNICRM_REDIS_LIST[name]
	if redis == nil || !ok {
		panic(fmt.Sprintf("未找到Redis配置: %s", name))
	}
	return redis
}

// 初始化时，将需要注册的表添加到RegisterTables中
func init() {
	RegisterTables = append(RegisterTables,
		system.SysApi{},
		system.JwtBlackList{},
	)
}
