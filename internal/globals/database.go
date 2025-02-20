package globals

import (
	"errors"

	"github.com/unicrm/server/internal/models/system"
	"gorm.io/gorm"
)

const (
	InitOrderSystem = iota
)

var (
	UNICRM_DB     *gorm.DB
	UNICRM_TABLES []interface{}
)

var (
	ErrMissingDBContext = errors.New("未找到数据库上下文")
	ErrDBTypeMismatch   = errors.New("数据库类型不匹配")
)

type InitContentKey string

// 初始化时，将需要注册的表添加到RegisterTables中
func init() {
	UNICRM_TABLES = append(UNICRM_TABLES,
		system.SysApi{},
		system.JwtBlackList{},
		system.SysUser{},
	)
}
