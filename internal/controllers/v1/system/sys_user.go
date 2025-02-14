package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/models/common/response"
	systemReq "github.com/unicrm/server/internal/models/system/request"
)

// Login 登录接口
// @Tags 系统管理
// @Summary 登录接口
// @Produce  application/json
// @Param user body systemReq.Login true "用户信息"
// @Success 200 {object} response.Response{}
// @Router /base/login [post]
func (baseApi *BaseApi) Login(c *gin.Context) {
	var user systemReq.Login
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ss := globals.UNICRM_CAPTCHA.Verify(user.CaptchaId, user.Captcha, true)
	fmt.Println("ssss:", ss)
	response.OkWithMessage("登录成功", c)
}
