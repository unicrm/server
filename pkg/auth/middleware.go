package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/pkg/auth/internal"
)

func (auth *AuthExtend) JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从cookie中获取x-token，如果获取不到，从header中获取，然后重新写入cookie中。
		token := auth.GetToken(c)
		if token == "" {
			internal.NoAuth("未登录或非法访问", c)
			c.Abort()
			return
		}

		// 检测是否在黑名单中
		if IsBlacklist(token) {
			internal.NoAuth("您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}

		// 解析token包含的信息
		claims, err := auth.ParseToken(token)
		if err != nil {
			if errors.Is(err, internal.ErrTokenExpired) {
				internal.NoAuth("令牌已过期，请重新登录", c)
				auth.ClearToken(c)
				c.Abort()
				return
			}
			internal.NoAuth("令牌解析失败", c)
			auth.ClearToken(c)
			c.Abort()
			return
		}

		// 刷新token，并重新写入cookie中
		newClaims, newtoken := auth.RefreshToken(claims, token)
		auth.SetToken(c, newtoken, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))

		// 将claims信息写入上下文
		c.Set("uuid", claims.UUID.String())
		c.Header("new-token", newtoken)
		c.Header("new-expires-at", fmt.Sprintf("%d", newClaims.ExpiresAt.Unix()))

		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}
