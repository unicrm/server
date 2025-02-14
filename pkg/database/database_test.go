package database

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/unicrm/server/pkg/database/internal"
	"github.com/unicrm/server/pkg/database/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	TEST_DB           *gorm.DB
	TEST_MYSQL_CONFIG *models.Mysql
)

func TestMain(m *testing.M) {
	// 初始化日志
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	// 初始化配置文件
	v := viper.New()
	v.SetConfigFile("../../config.debug.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logger.Error("读取配置文件失败", zap.Error(err))
	}
	v.WatchConfig()
	if err := v.UnmarshalKey("mysql", &TEST_MYSQL_CONFIG); err != nil {
		logger.Error("解析配置文件失败", zap.Error(err))
	}
	m.Run()
}

func TestDatabaseMysql(t *testing.T) {
	general := models.GeneralDB(TEST_MYSQL_CONFIG.GeneralDB)
	TEST_DB = internal.MysqlInitByConfig(general)
	assert.NotNil(t, TEST_DB)
}

func TestTables(t *testing.T) {
	registerTables := []interface{}{
		models.Test{},
	}
	RegisterTables(TEST_DB, registerTables...)
}
