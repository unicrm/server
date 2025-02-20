package request

// 注册请求结构体
type Register struct {
	Username string `json:"username" validate:"required" example:"用户名"` // 用户名
	Password string `json:"password" validate:"required" example:"密码"`  // 密码
	Enable   int    `json:"enable" swaggertype:"string" example:"是否启用"` // 是否启用
	Phone    string `json:"phone" example:"手机号"`                        // 手机号
	Email    string `json:"email" example:"邮箱"`                         // 邮箱
}

// 登录请求结构体
type Login struct {
	Username  string `json:"username" validate:"required"` // 用户名
	Password  string `json:"password" validate:"required"` // 密码
	Captcha   string `json:"captcha" validate:"required"`  // 验证码
	CaptchaId string `json:"captcha_id"`                   // 验证码ID
}
