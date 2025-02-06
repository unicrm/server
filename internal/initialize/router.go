package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/unicrm/server/docs"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/middleware"
)

// Routers 初始化路由信息
func Routers() *gin.Engine {

	Router := gin.Default()

	// 注册swagger文档路由
	docs.SwaggerInfo.BasePath = globals.UNICRM_CONFIG.System.RouterPrefix
	Router.GET(globals.UNICRM_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	globals.UNICRM_LOGGER.Info("注册文档路由成功")

	AllowAnyGroup := Router.Group(globals.UNICRM_CONFIG.System.RouterPrefix)
	IsAuthenticatedGroup := Router.Group(globals.UNICRM_CONFIG.System.RouterPrefix)
	IsAdminUser := Router.Group(globals.UNICRM_CONFIG.System.RouterPrefix)

	IsAuthenticatedGroup.Use(middleware.JWTAuthMiddleware())
	IsAdminUser.Use()

	{
		// 健康检查接口
		AllowAnyGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	return Router

}
