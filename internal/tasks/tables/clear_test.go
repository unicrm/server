package tables

import (
	"os"
	"testing"

	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/initialize"
	"github.com/unicrm/server/pkg/database"
	"go.uber.org/zap"
)

func TestMain(m *testing.M) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	exit := m.Run()
	os.Exit(exit)
}

func TestDatabaseInit(t *testing.T) {
	// 初始化数据库
	initialize.LoadConfig("../../../config.debug.yaml")
	globals.UNICRM_DB = database.DatabaseInit(globals.UNICRM_CONFIG.GeneralDB)
}

func TestClearTable(t *testing.T) {
	err := ClearTable(globals.UNICRM_DB, "unicrm_jwt_black_list")
	zap.L().Error(err.Error())
}

func TestEnd(t *testing.T) {
	defer database.DatabaseClose(globals.UNICRM_DB)
}
