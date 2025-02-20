package system

import (
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/common/response"
	"github.com/unicrm/server/internal/models/system/request"
	"github.com/unicrm/server/internal/services/system"
	"go.uber.org/zap"
)

type InitApiInterface interface {
	InitDB(c *gin.Context)
	CheckDB(c *gin.Context)
}

type InitApi struct{}

var InitApiApp InitApiInterface = new(InitApi)

// InitDB
// @Summary 初始化数据库
// @Tags 数据库
// @Produce  application/json
// @Param    data  body      request.InitDB                  true  "初始化数据库参数"
// @Success  200   {object}  response.Response{data=string}  "初始化用户数据库"
// @Router /init/initdb [post]
func (initApi *InitApi) InitDB(c *gin.Context) {
	if globals.UNICRM_DB != nil {
		zap.L().Error("数据库已初始化")
		response.FailWithMessage("数据库已初始化", c)
		return
	}
	var initDB request.InitDB
	if err := c.ShouldBind(&initDB); err != nil {
		zap.L().Error("参数错误", zap.Any("err", err))
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := system.InitDBServiceApp.InitDB(initDB); err != nil {
		zap.L().Error("初始化数据库失败", zap.Error(err))
		response.FailWithMessage("初始化数据库失败", c)
		return
	}
	response.OkWithMessage("初始化数据库成功", c)
}

// CheckDB
// @Summary 检查数据库
// @Tags 数据库
// @Produce  application/json
// @Success  200   {object}  response.Response{data=map[string]interface{}, msg=string}  "检查数据库是否初始化成功"
// @Router /init/checkdb [post]
func (initApi *InitApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if globals.UNICRM_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	zap.L().Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
}
