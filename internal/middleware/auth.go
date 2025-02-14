package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/services/system"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 如果开启了redis，则将token和uuid存入redis
		if globals.UNICRM_CONFIG.System.UseRedis {
			if uuid, exists := c.Get("uuid"); exists {
				token, _ := c.Cookie("x-token")
				system.AuthServiceApp.SetRedis(token, uuid.(string))
			}
		}

		// 继续执行后续的中间件
		c.Next()
	}
}
