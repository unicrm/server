package system

import "github.com/gin-gonic/gin"

type BaseApiInterface interface {
	Captcha(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type BaseApi struct{}

var BaseApiApp BaseApiInterface = new(BaseApi)
