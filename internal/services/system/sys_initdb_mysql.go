package system

import (
	"context"

	"github.com/unicrm/server/internal/models/system/request"
)

type MysqlInitHandler struct{}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// WriteConfig mysql回写配置
func (h MysqlInitHandler) WriteConfig(ctx context.Context) error {

	return nil
}

// EnsureDB 创建数据库并初始化 mysql
func (h MysqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {

	return next, err
}
