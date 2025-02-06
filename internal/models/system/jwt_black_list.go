package system

import (
	"github.com/unicrm/server/internal/config"
)

type JwtBlackList struct {
	config.UNICRM_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
