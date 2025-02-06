package initialize

import (
	// 初始化系统数据
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/unicrm/server/internal/globals"
	_ "github.com/unicrm/server/internal/source/system"
)

func OtherInit() {
	jwt := globals.UNICRM_CONFIG.JWT
	fmt.Printf("加载其他配置项, JWT: %+v \n", jwt)

	dr, _ := time.ParseDuration(jwt.BufferTime)
	globals.UNICRM_BLACK_CACHE = cache.New(dr, 10*time.Minute)
}

func init() {}
