package internal

import "errors"

var (
	ErrTokenParsed      = errors.New("令牌解析失败")
	ErrTokenExpired     = errors.New("令牌已过期")
	ErrTokenNotValidYet = errors.New("令牌尚未生效")
	ErrTokenMalformed   = errors.New("错误的令牌格式")
	ErrTokenInvalid     = errors.New("无效的令牌")
	ErrClaimsInvalid    = errors.New("无效的声明")
)
