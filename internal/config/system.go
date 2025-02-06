package config

type System struct {
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	DBType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
