package auth

import (
	"errors"

	"github.com/google/uuid"
	"github.com/unicrm/server/pkg/auth/internal"
)

// 登录认证生成token
func LoginToken(uuid uuid.UUID) (token string, claims internal.CustomClaims, err error) {
	if AUTH == nil {
		return "", claims, errors.New("认证服务尚未初始化")
	}
	claims = AUTH.CreateClaims(internal.BaseClaims{
		UUID: uuid,
	})
	token, err = AUTH.CreateToken(claims)
	return token, claims, err
}
