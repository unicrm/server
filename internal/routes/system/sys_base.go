package system

import (
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/controllers/v1/system"
)

var BaseRouterApp = new(BaseRouter)

type BaseRouter struct{}

func (BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("register", system.BaseApiApp.Register)
		baseRouter.POST("login", system.BaseApiApp.Login)
		baseRouter.POST("captcha", system.BaseApiApp.Captcha)
	}
	return baseRouter
}
