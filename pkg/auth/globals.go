package auth

import (
	"github.com/patrickmn/go-cache"
	"golang.org/x/sync/singleflight"
)

var (
	// 全局鉴权变量
	AUTH *AuthExtend
	// 黑名单缓存
	AUTH_BLACK_CACHE *cache.Cache
	// 单例模式缓存锁
	AUTH_SINGLE_FLIGHT *singleflight.Group = &singleflight.Group{}
)

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}

var _ AuthExtendInterface = new(AuthExtend)

type AuthExtend struct {
	JWT
	SigningKeyByte []byte
}

// 初始化鉴权配置
func InitAuth(jwtConfig JWT) *AuthExtend {
	AUTH = &AuthExtend{
		JWT:            jwtConfig,
		SigningKeyByte: []byte(jwtConfig.SigningKey),
	}
	return AUTH
}
