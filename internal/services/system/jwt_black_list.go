package system

import (
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/system"
	"go.uber.org/zap"
)

func LoadJwtBlackList() {
	if globals.UNICRM_DB == nil {
		globals.UNICRM_LOGGER.Error("加载数据库jwt黑名单失败, 原因: 全局数据库未初始化")
		return
	}
	var datas []string
	err := globals.UNICRM_DB.Model(system.JwtBlackList{}).Select("jwt").Find(&datas).Error
	if err != nil {
		globals.UNICRM_LOGGER.Error("加载数据库jwt黑名单失败, 原因: 数据库查询失败", zap.Error(err))
		return
	}
	if len(datas) == 0 {
		globals.UNICRM_LOGGER.Warn("加载数据库jwt黑名单成功, 但未查询到数据")
		return
	}
	for _, data := range datas {
		globals.UNICRM_LOGGER.Info("加载数据库jwt黑名单成功, 已添加到缓存", zap.String("jwt", data))
		globals.UNICRM_BLACK_CACHE.SetDefault(data, struct{}{})
	}
}
