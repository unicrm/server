package request

import (
	"github.com/unicrm/server/pkg/database/model"
)

type InitDB struct {
	model.GeneralDB
	AdminPassword string `json:"admin-password" binding:"required"`
	DBName        string `json:"db-name" binding:"required"` // 数据库名
	DBType        string `json:"db-type"`                    // 数据库类型
}

func (i InitDB) ToMysqlConfig() model.Mysql {
	return model.Mysql{
		GeneralDB: model.GeneralDB(i.GeneralDB),
	}
}
