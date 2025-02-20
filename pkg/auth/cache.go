package auth

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/unicrm/server/pkg/auth/tools"
	"go.uber.org/zap"
)

// 从本地缓存中判断jwt是否在黑名单中
func IsBlacklist(jwt string) bool {
	if AUTH_BLACK_CACHE == nil {
		zap.L().Error("判断jwt是否在黑名单失败, 原因: 全局缓存未初始化")
		return true
	}
	_, ok := AUTH_BLACK_CACHE.Get(jwt)
	return ok
}

// 将jwt加入黑名单
func AddBlacklist(jwt string) error {
	if AUTH_BLACK_CACHE == nil {
		zap.L().Error("将jwt加入黑名单失败, 原因: 全局缓存未初始化")
		return errors.New("全局缓存未初始化")
	}
	AUTH_BLACK_CACHE.SetDefault(jwt, true)
	return nil
}

// 初始化jwt黑名单缓存
func InitBlacklist(jwt JWT) (abc *cache.Cache, err error) {
	dr, err := tools.ParseDuration(jwt.ExpiresTime)
	if err != nil {
		return nil, err
	}
	AUTH_BLACK_CACHE = cache.New(dr, 10*time.Minute)
	return AUTH_BLACK_CACHE, nil
}

// 初始化jwt到黑名单
func SetBlacklist(datas []string) error {
	if AUTH_BLACK_CACHE == nil {
		zap.L().Error("将jwt加入黑名单失败, 原因: 全局缓存未初始化")
		return errors.New("全局缓存未初始化")
	}
	for _, data := range datas {
		AUTH_BLACK_CACHE.SetDefault(data, struct{}{})
	}
	return nil
}
