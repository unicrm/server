package system

import (
	"context"

	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/pkg/auth/tools"
)

type AuthServiceInterface interface {
	SetRedis(token string, user string) error
}

type AuthService struct{}

var AuthServiceApp AuthServiceInterface = new(AuthService)

func (authService *AuthService) GetRedis(user string) (token string, err error) {
	return globals.UNICRM_REDIS.Get(context.Background(), user).Result()
}

// SetRedis 设置redis缓存
func (authService *AuthService) SetRedis(token string, user string) (err error) {
	dr, err := tools.ParseDuration(globals.UNICRM_AUTH.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = globals.UNICRM_REDIS.Set(context.Background(), user, token, timer).Err()
	return err
}
