package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unicrm/server/pkg/database/internal"
	"github.com/unicrm/server/pkg/database/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var TEST_DB *gorm.DB

func TestDatabaseMysql(t *testing.T) {
	general := model.GeneralDB{
		DBType:       internal.DatabaseDBType,
		Host:         internal.DatabaseHost,
		Port:         internal.DatabasePort,
		User:         internal.DatabaseUser,
		Password:     internal.DatabasePassword,
		DBName:       internal.DatabaseDBName,
		Prefix:       internal.DatabasePrefix,
		Singular:     internal.DatabaseSingular,
		Config:       internal.DatabaseConfig,
		MaxIdleConns: internal.DatabaseMaxIdleConns,
		MaxOpenConns: internal.DatabaseMaxOpenConns,
		Engine:       internal.DatabaseEngine,
		LogMode:      internal.DatabaseLogMode,
	}
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	TEST_DB = internal.MysqlInitByConfig(general)
	assert.NotNil(t, TEST_DB)
}

func TestTables(t *testing.T) {
	registerTables := []interface{}{
		model.Test{},
	}
	RegisterTables(TEST_DB, registerTables...)
}
