package request

import (
	"github.com/unicrm/server/pkg/database/models"
)

type InitDB struct {
	models.GeneralDB
	AdminPassword string `json:"admin-password" binding:"required"`
	DBName        string `json:"db-name" binding:"required"` // 数据库名
	DBType        string `json:"db-type"`                    // 数据库类型
}

func (i InitDB) ToMysqlConfig() models.Mysql {
	return models.Mysql{
		GeneralDB: models.GeneralDB(i.GeneralDB),
	}
}
