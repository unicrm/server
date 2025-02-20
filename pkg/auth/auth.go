package auth

import (
	"net"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/unicrm/server/pkg/auth/internal"
	"github.com/unicrm/server/pkg/auth/tools"
	"go.uber.org/zap"
)

// 创建自定义声明
func (auth *AuthExtend) CreateClaims(baseClaims BaseClaims) CustomClaims {
	bf, _ := tools.ParseDuration(auth.BufferTime)
	ep, _ := tools.ParseDuration(auth.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"UNICRM_AUTH"},        // 受众
			NotBefore: jwt.NewNumericDate(time.Now()),         // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间 7天  配置文件
			Issuer:    auth.Issuer,                            // 签名的发行者
		},
	}
	return claims
}

// 创建令牌
func (auth *AuthExtend) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AUTH.SigningKeyByte)
}

// 解析令牌
func (auth *AuthExtend) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return AUTH.SigningKeyByte, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		zap.L().Error(internal.ErrTokenParsed.Error(), zap.Error(err))
		return nil, err
	} else if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	} else {
		zap.L().Error(internal.ErrClaimsInvalid.Error())
		return nil, internal.ErrClaimsInvalid
	}
}

// 获取令牌
func (auth *AuthExtend) GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
		claims, err := auth.ParseToken(token)
		if err != nil {
			zap.L().Error("重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构", zap.Error(err))
			return token
		}
		if claims == nil {
			zap.L().Error("重新写入cookie token失败,claims为空", zap.Error(err))
			return token
		}
		auth.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	}
	return token
}

// 设置令牌
func (auth *AuthExtend) SetToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", host, false, true)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", "", false, true)
	}
}

// 清除令牌
func (auth *AuthExtend) ClearToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", host, false, true)
	} else {
		c.SetCookie("x-token", "", -1, "/", "", false, true)
	}
}

// 刷新令牌
func (auth *AuthExtend) RefreshToken(claims *CustomClaims, token string) (*CustomClaims, string) {
	if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
		dr, _ := tools.ParseDuration(auth.ExpiresTime)
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
		// 使用归并回源避免并发问题
		v, err, _ := AUTH_SINGLE_FLIGHT.Do("AUTH:"+token, func() (interface{}, error) {
			return auth.CreateToken(*claims)
		})
		if err != nil {
			zap.L().Error("刷新令牌失败", zap.Error(err))
			// 返回旧令牌
			return claims, token
		}
		return claims, v.(string)
	}
	return claims, token
}
