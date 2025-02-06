package system

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/system/request"
)

// SubInitializer 提供 source/*/init() 使用的接口，每个 initializer 完成一个初始化过程
type SubInitializer interface {
	InitializerName() string
	MigrateTable(ctx context.Context) (context.Context, error)
	TableCreate(ctx context.Context) bool
	InitializeData(ctx context.Context) (context.Context, error)
}

// TypedDBInitHandler 执行传入的 initializer
type TypedDBInitHandler interface {
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error) // 建库，失败属于 fatal error，因此让它 panic
	WriteConfig(ctx context.Context) error                                       // 回写配置
	// InitTables(ctx context.Context, inits initSlice) error                       // 建表 handler
	// InitData(ctx context.Context, inits initSlice) error                         // 建数据 handler
}

type orderedInitializer struct {
	order int
	SubInitializer
}

type initSlice []*orderedInitializer

var (
	initializers = initSlice{}
	cache        = map[string]*orderedInitializer{}
)

// RegisterInit 注册要执行的初始化过程，会在 InitDB() 时调用
func RegisterInit(order int, init SubInitializer) {
	name := init.InitializerName()
	if _, existed := cache[name]; existed {
		panic(fmt.Sprintf("重复注册初始化器 %s", name))
	}
	ni := orderedInitializer{order, init}
	initializers = append(initializers, &ni)
	cache[name] = &ni
}

/* ---- * service * ---- */

type InitDBService struct{}

func (i *InitDBService) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, globals.InitContentKey("admin-password"), conf.AdminPassword)
	if len(initializers) == 0 {
		return errors.New("没有初始化器，请检查配置文件是否有误")
	}
	sort.Sort(initializers)

	var initHandler TypedDBInitHandler
	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, globals.InitContentKey("dbtype"), "mysql")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, globals.InitContentKey("dbtype"), "mysql")
	}
	ctx, err = initHandler.EnsureDB(ctx, &conf)
	if err != nil {
		return err
	}
	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	initializers = initSlice{}               // 清空 initializers
	cache = map[string]*orderedInitializer{} // 清空 cache
	return nil
}

/* -- sortable interface -- */

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
