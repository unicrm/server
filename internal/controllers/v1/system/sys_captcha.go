package system

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/common/response"
	systemRes "github.com/unicrm/server/internal/models/system/response"
	"github.com/unicrm/server/pkg/auth"
	"go.uber.org/zap"
)

// Captcha 获取验证码
// @Tags 系统管理
// @Summary 获取验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/json
// @Success 200 {object} response.Response{data=systemRes.SysCaptchaResponse}
// @Router /base/captcha [post]
func (baseApi *BaseApi) Captcha(c *gin.Context) {

	if globals.UNICRM_CAPTCHA == nil {
		zap.L().Error("验证码获取失败", zap.Any("error", errors.New("未初始化验证码配置信息")))
		response.OkWithMessage("验证码获取失败", c)
		return
	}

	// 验证码配置信息
	config := globals.UNICRM_CONFIG.Captcha

	// 判断验证码是否开启
	key := c.ClientIP()
	v, ok := auth.AUTH_BLACK_CACHE.Get(key)
	if !ok {
		auth.AUTH_BLACK_CACHE.Set(key, 1, time.Second*time.Duration(config.OpenCaptchaTimeout))
		v = 1
	}

	captchaId, image, _, err := globals.UNICRM_CAPTCHA.Generate()
	if err != nil {
		zap.L().Error("验证码获取失败", zap.Any("error", err))
		response.OkWithMessage("验证码获取失败", c)
		return
	}
	// 返回客户端需要的数据格式
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     captchaId,
		PicPath:       image,
		CaptchaLength: config.KeyLong,
		OpenCaptcha:   config.OpenCaptcha == 0 || config.OpenCaptcha < v.(int),
	}, "验证码获取成功", c)
}
