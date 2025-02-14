package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/unicrm/server/docs"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/middleware"
	"github.com/unicrm/server/internal/routes/system"
	"go.uber.org/zap"
)

// Routers 初始化路由信息
func Routers() *gin.Engine {

	gin.SetMode(globals.UNICRM_CONFIG.System.RunMode)

	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(gin.Logger())

	// 注册swagger文档路由
	docs.SwaggerInfo.BasePath = globals.UNICRM_CONFIG.System.RouterPrefix
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	zap.L().Info("注册文档路由成功")

	AllowAnyGroup := Router.Group(globals.UNICRM_CONFIG.System.RouterPrefix)
	IsAuthenticatedGroup := Router.Group(globals.UNICRM_CONFIG.System.RouterPrefix)
	IsAdminUser := Router.Group(globals.UNICRM_CONFIG.System.RouterPrefix)

	IsAuthenticatedGroup.Use(globals.UNICRM_AUTH.JWTAuthMiddleware(), middleware.JWTAuthMiddleware())
	IsAdminUser.Use()

	{
		// 健康检查接口
		AllowAnyGroup.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	{
		// jwt测试
		IsAuthenticatedGroup.POST("/jwt", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})
	}

	{
		system.BaseRouterApp.InitBaseRouter(AllowAnyGroup)
	}
	zap.L().Info("注册路由成功")
	return Router

}
