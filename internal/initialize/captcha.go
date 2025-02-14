package initialize

import (
	"github.com/mojocn/base64Captcha"
	"github.com/unicrm/server/internal/globals"
)

func CaptchaInit() *base64Captcha.Captcha {
	// 验证码配置信息
	config := globals.UNICRM_CONFIG.Captcha
	// 配置验证码的参数
	store := base64Captcha.DefaultMemStore
	// 生成验证码驱动，这里使用数字类型验证码
	driver := base64Captcha.NewDriverDigit(config.ImageHeight, config.ImageWidth, config.KeyLong, 0.7, 80)
	// 创建一个验证码生成器
	return base64Captcha.NewCaptcha(driver, store)
}
