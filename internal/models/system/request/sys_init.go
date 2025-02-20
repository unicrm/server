package request

import (
	"fmt"

	"github.com/unicrm/server/pkg/database/models"
)

type InitDB struct {
	models.GeneralDB
	AdminPassword string `json:"admin-password" binding:"required"`
	DBName        string `json:"db-name" binding:"required"` // 数据库名
	DBType        string `json:"db-type"`                    // 数据库类型
}

// MysqlEmptyDsn 空数据库
func (initDB InitDB) MysqlEmptyDsn() string {
	if initDB.Host == "" {
		initDB.Host = "127.0.0.1"
	}
	if initDB.Port == "" {
		initDB.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", initDB.User, initDB.Password, initDB.Host, initDB.Port)
}

// ToMysqlConfig 转成mysql配置
func (initDB InitDB) ToMysqlConfig() models.Mysql {
	return models.Mysql{
		GeneralDB: models.GeneralDB(initDB.GeneralDB),
	}
}
