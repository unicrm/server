package auth

import (
	"errors"

	"github.com/google/uuid"
)

// 登录认证生成token
func LoginToken(uuid uuid.UUID) (token string, claims CustomClaims, err error) {
	if AUTH == nil {
		return "", claims, errors.New("认证服务尚未初始化")
	}
	claims = AUTH.CreateClaims(BaseClaims{
		UUID: uuid,
	})
	token, err = AUTH.CreateToken(claims)
	return token, claims, err
}
