package models

import (
	"strings"

	"gorm.io/gorm/logger"
)

type GeneralDB struct {
	DBType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	DBName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // 数据库引擎，默认InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
}

func (generalDB GeneralDB) LogLevel() logger.LogLevel {
	switch strings.ToLower(generalDB.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
