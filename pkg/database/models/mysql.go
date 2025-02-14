package models

import (
	"net/url"
	"strings"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// DSNConfig 返回mysql.Config对象
func (db *Mysql) DSNConfig() *mysql.Config {
	// 解析配置文件中的参数
	paramsValues, err := url.ParseQuery(db.Config)
	if err != nil {
		zap.L().Panic("解析数据库配置出错", zap.Error(err))
	}
	// 将paramsValues转换为map[string]string类型
	params := make(map[string]string, len(paramsValues))
	for k, v := range paramsValues {
		params[k] = strings.Join(v, "")
	}
	// 创建mysql.Config对象
	return &mysql.Config{
		User:                 db.User,
		Passwd:               db.Password,
		Net:                  "tcp",
		Addr:                 db.Host + ":" + db.Port,
		DBName:               db.DBName,
		Params:               params,
		AllowNativePasswords: true,
	}
}
