package system

import (
	"github.com/unicrm/server/internal/config"
)

type SysApi struct {
	config.UNICRM_MODEL
	Path        string `json:"path" gorm:"comment:接口路径"`
	Description string `json:"description" gorm:"comment:接口描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:接口分组"`
	Method      string `json:"method" gorm:"default:POST;comment:请求方式"`
}

func (SysApi) TableName() string {
	return "sys_api"
}
