package system

import (
	"context"
	"errors"

	"github.com/unicrm/server/internal/globals"
	reqModel "github.com/unicrm/server/internal/models/system"
	"github.com/unicrm/server/internal/services/system"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = globals.InitOrderSystem + 1

func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i *initApi) InitializerName() string {
	return reqModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value(globals.InitContentKey("db")).(*gorm.DB)
	if !ok {
		return ctx, globals.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&reqModel.SysApi{})
}

func (i *initApi) TableCreate(ctx context.Context) bool {
	db, ok := ctx.Value(globals.InitContentKey("db")).(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&reqModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value(globals.InitContentKey("db")).(*gorm.DB)
	if !ok {
		return ctx, globals.ErrMissingDBContext
	}
	entities := []reqModel.SysApi{
		{ApiGroup: "api", Method: "POST", Path: "/api/create", Description: "创建API接口"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.New("初始化数据失败")
	}
	next := context.WithValue(ctx, globals.InitContentKey(i.InitializerName()), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value(globals.InitContentKey("db")).(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ?", "/api/createApi").
		First(&reqModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
