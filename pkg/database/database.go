package database

import (
	"github.com/unicrm/server/pkg/database/internal"
	"github.com/unicrm/server/pkg/database/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func DatabaseInit(general model.GeneralDB) *gorm.DB {
	switch general.DBType {
	case "mysql":
		return internal.MysqlInitByConfig(general)
	default:
		return internal.MysqlInitByConfig(general)
	}
}

func DatabaseClose(db *gorm.DB) {
	if db == nil {
		zap.L().Warn("数据库关闭失败, 原因: 未传入数据库连接")
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("数据库关闭失败", zap.Error(err))
		return
	}
	if err := sqlDB.Close(); err != nil {
		zap.L().Error("数据库关闭失败", zap.Error(err))
		return
	}
	zap.L().Info("数据库已成功关闭")
}

func RegisterTables(db *gorm.DB, dst ...interface{}) {
	if db == nil {
		zap.L().Fatal("数据库表注册失败, 原因: 未传入数据库连接")
		return
	}
	if len(dst) == 0 {
		zap.L().Warn("数据库表注册失败, 原因: 未传入任何模型")
		return
	}
	err := db.AutoMigrate(dst...)
	if err != nil {
		zap.L().Fatal("数据库表注册失败, 原因: 数据库迁移失败", zap.Error(err))
		return
	}
	zap.L().Info("数据库表注册成功")
}
