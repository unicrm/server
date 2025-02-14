package config

import (
	"github.com/unicrm/server/pkg/auth"
	"github.com/unicrm/server/pkg/database/models"
	"github.com/unicrm/server/pkg/logging"
	"github.com/unicrm/server/pkg/redis"
)

type UnicrmConfig struct {
	Logger    logging.Logger   `mapstructure:"logger" json:"logger" yaml:"logger"`
	System    System           `mapstructure:"system" json:"system" yaml:"system"`
	JWT       auth.JWT         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	GeneralDB models.GeneralDB `mapstructure:"general-db" json:"general-db" yaml:"general-db"`
	Redis     redis.Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []redis.Redis    `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	Captcha   Captcha          `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
