package system

import (
	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/common/response"
	"github.com/unicrm/server/internal/models/system"
	systemReq "github.com/unicrm/server/internal/models/system/request"
	systemRes "github.com/unicrm/server/internal/models/system/response"
	systemService "github.com/unicrm/server/internal/services/system"
	"go.uber.org/zap"
)

// Register
// @Summary 注册接口
// @Tags 系统管理
// @Produce  application/json
// @Param user body systemReq.Register true "用户信息"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse{user=system.SysUser}}
// @Router /base/register [post]
func (baseApi *BaseApi) Register(c *gin.Context) {
	var register systemReq.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := globals.UNICRM_VALIDATE.Struct(register); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userReturn, err := systemService.UserServiceApp.Register(&system.SysUser{
		Username: register.Username,
		Password: register.Password,
		Phone:    register.Phone,
		Email:    register.Email,
		Enable:   register.Enable,
	})
	if err != nil {
		zap.L().Error("注册失败", zap.Any("error", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// Login
// @Summary 登录接口
// @Tags 系统管理
// @Produce  application/json
// @Param user body systemReq.Login true "用户信息"
// @Success 200 {object} response.Response{}
// @Router /base/login [post]
func (baseApi *BaseApi) Login(c *gin.Context) {
	var login systemReq.Login
	key := c.ClientIP()
	if err := c.ShouldBindJSON(&login); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 参数校验
	if err := globals.UNICRM_VALIDATE.Struct(login); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证码校验
	// if !globals.UNICRM_CAPTCHA.Verify(login.CaptchaId, login.Captcha, true) {
	// 	globals.UNICRM_BLACK_CACHE.Increment(key, 1)
	// 	response.FailWithMessage("验证码错误", c)
	// }
	user := &system.SysUser{Username: login.Username, Password: login.Password}
	userReturn, err := systemService.UserServiceApp.Login(user)
	if err != nil {
		zap.L().Error("用户名不存在或者密码错误", zap.Any("error", err))
		// 验证码次数+1
		globals.UNICRM_BLACK_CACHE.Increment(key, 1)
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if userReturn.Enable != 1 {
		zap.L().Error("用户未启用", zap.Any("error", err))
		// 验证码次数+1
		globals.UNICRM_BLACK_CACHE.Increment(key, 1)
		response.FailWithMessage("用户未启用", c)
		return
	}
	token, msg, err := systemService.AuthServiceApp.CreateToken(c, *userReturn)
	if err != nil {
		zap.L().Error(msg, zap.Any("error", err))
		// 验证码次数+1
		globals.UNICRM_BLACK_CACHE.Increment(key, 1)
		response.FailWithMessage(msg, c)
		return
	}
	response.OkWithDetailed(systemRes.LoginResponse{
		User:  *userReturn,
		Token: token,
	}, "登录成功", c)
}
