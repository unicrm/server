package config

type Captcha struct {
	ImageWidth         int `mapstructure:"image-width" json:"image-width" yaml:"image-width"`
	ImageHeight        int `mapstructure:"image-height" json:"image-height" yaml:"image-height"`
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`
	OpenCaptchaTimeout int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"`
}
