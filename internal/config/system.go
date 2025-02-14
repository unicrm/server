package config

type System struct {
	RunMode      string `mapstructure:"run-mode" json:"run-mode" yaml:"run-mode"`
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	RedisName    string `mapstructure:"redis-name" json:"redis-name" yaml:"redis-name"`
}
