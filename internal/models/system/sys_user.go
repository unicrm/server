package system

import (
	"github.com/google/uuid"
	"github.com/unicrm/server/internal/config"
)

type Login interface {
	Login(username, password string) (bool, error)
	GetUsername() string
}

var _ Login = new(SysUser)

type SysUser struct {
	config.UNICRM_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户uuid"`
	Username string    `json:"username" gorm:"index;comment:用户名"`
	Password string    `json:"-" gorm:"comment:密码"`
	Phone    string    `json:"phone" gorm:"comment:手机号"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Enable   int       `json:"enable" gorm:"default:1;comment:是否启用 1启用 0禁用"`
}

func (s *SysUser) Login(username, password string) (bool, error) {
	return true, nil
}

func (s *SysUser) GetUsername() string {
	return s.Username
}
