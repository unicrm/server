package system

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/system"
	"github.com/unicrm/server/pkg/auth"
	"github.com/unicrm/server/pkg/auth/tools"
)

type AuthServiceInterface interface {
	JsonInBlackCache(jwtList system.JwtBlackList) (err error)
	GetRedis(user string) (token string, err error)
	SetRedis(token string, user string) error
	CreateToken(c *gin.Context, user system.SysUser) (token string, msg string, err error)
}

type AuthService struct{}

var AuthServiceApp AuthServiceInterface = new(AuthService)

func (authService *AuthService) JsonInBlackCache(jwtList system.JwtBlackList) (err error) {
	err = globals.UNICRM_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	globals.UNICRM_BLACK_CACHE.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (authService *AuthService) GetRedis(userName string) (token string, err error) {
	redisToken := globals.UNICRM_REDIS.Get(context.Background(), userName).Val()
	return redisToken, nil
}

// SetRedis 设置redis缓存
func (authService *AuthService) SetRedis(token string, userName string) (err error) {
	dr, err := tools.ParseDuration(globals.UNICRM_AUTH.ExpiresTime)
	if err != nil {
		return err
	}
	err = globals.UNICRM_REDIS.Set(context.Background(), userName, token, dr).Err()
	return err
}

// 签发token
func (authService *AuthService) CreateToken(c *gin.Context, user system.SysUser) (token string, msg string, err error) {
	// 生成token
	newToken, newClaims, err := auth.LoginToken(user.UUID)
	if err != nil {
		return "", "获取token失败", err
	}
	// 多点登录，直接返回新的token
	if globals.UNICRM_CONFIG.System.UseMultipoint {
		// 设置token到cookie中，并设置过期时间
		globals.UNICRM_AUTH.SetToken(c, newToken, int(newClaims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		// 返回新的token
		return newToken, "签发token成功", nil
	}
	// 因为是单点登录，所以如果redis中有该用户的token，则加入黑名单缓存
	// 获取redis中的token
	if redisToken, err := authService.GetRedis(user.Username); err != nil {
		// 获取redis中的token失败
		return "", "设置登录状态失败", err
	} else if redisToken != "" {
		// 如果redis中存在该用户token，则加入黑名单缓存
		if redisToken != "" {
			blackCache := system.JwtBlackList{Jwt: redisToken}
			if err := authService.JsonInBlackCache(blackCache); err != nil {
				return "", "设置登录状态失败", err
			}
		}
	}
	// 向redis中设置新的token
	if err := authService.SetRedis(newToken, user.Username); err != nil {
		return "", "设置登录状态失败", err
	}
	// 设置token到cookie中，并设置过期时间
	globals.UNICRM_AUTH.SetToken(c, newToken, int(newClaims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	// 返回新的token
	return newToken, "签发token成功", nil
}
