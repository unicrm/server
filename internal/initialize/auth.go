package initialize

import (
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/system"
	"github.com/unicrm/server/pkg/auth"
	"go.uber.org/zap"
)

func AuthInit() *auth.AuthExtend {
	// 初始化JWT认证服务
	authExtend := auth.InitAuth(globals.UNICRM_CONFIG.JWT)
	// 初始化，并设置缓存过期时间
	globals.UNICRM_BLACK_CACHE, _ = auth.InitBlacklist(globals.UNICRM_CONFIG.JWT)
	return authExtend
}

// 从数据库加载jwt黑名单到本地缓存
func LoadJwtBlackList() (datas []string) {
	if globals.UNICRM_DB == nil {
		zap.L().Error("加载数据库jwt黑名单失败, 原因: 全局数据库未初始化")
		return
	}
	err := globals.UNICRM_DB.Model(system.JwtBlackList{}).Select("jwt").Find(&datas).Error
	if err != nil {
		zap.L().Error("加载数据库jwt黑名单失败, 原因: 数据库查询失败", zap.Error(err))
		return
	}
	auth.SetBlacklist(datas)
	zap.L().Info("加载数据库jwt黑名单成功, 已添加到缓存", zap.Strings("jwts", datas))
	return
}
