package internal

import (
	"github.com/unicrm/server/pkg/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 自定义配置初始化数据库连接
func MysqlInitByConfig(general models.GeneralDB) *gorm.DB {
	database := models.Mysql{GeneralDB: general}
	mysqlConfig := mysql.Config{
		DSNConfig:                 database.DSNConfig(), // DSN data source name
		DefaultStringSize:         256,                  // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config(general)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+database.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(database.MaxIdleConns)
		sqlDB.SetMaxOpenConns(database.MaxOpenConns)
		return db
	}
}
