package system

import (
	"context"

	"github.com/unicrm/server/internal/globals"
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
	// 校验数据库类型是否为mysql
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, globals.ErrDBTypeMismatch
	}
	// 转化成mysql配置
	c := conf.ToMysqlConfig()
	next = context.WithValue(ctx, globals.InitContentKey("config"), c)
	// 如果数据库名称为空，跳过初始化
	if c.DBName == "" {
		return ctx, nil
	}
	// 创建数据库
	dsn := conf.MysqlEmptyDsn()
	createSql := ""
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		return nil, err
	}
	c.DSNConfig()
	return next, err
}
