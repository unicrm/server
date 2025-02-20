package system

import (
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/controllers/v1/system"
)

var InitRouterApp = new(InitRouter)

type InitRouter struct{}

func (InitRouter) InitInitRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	initRouter := Router.Group("init")
	{
		initRouter.POST("initdb", system.InitApiApp.InitDB)
		initRouter.POST("checkdb", system.InitApiApp.CheckDB)
	}
	return initRouter
}
