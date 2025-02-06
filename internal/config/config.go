package config

import (
	"time"

	"github.com/unicrm/server/pkg/database/model"
	"github.com/unicrm/server/pkg/logging"
	"github.com/unicrm/server/pkg/redis"
	"gorm.io/gorm"
)

type UnicrmConfig struct {
	Logger    logging.Logger `mapstructure:"logger" json:"logger" yaml:"logger"`
	System    System         `mapstructure:"system" json:"system" yaml:"system"`
	JWT       JWT            `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql     model.Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis     redis.Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []redis.Redis  `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
}

type UNICRM_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
