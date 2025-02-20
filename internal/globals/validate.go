package globals

import "github.com/go-playground/validator/v10"

var (
	UNICRM_VALIDATE *validator.Validate
)

func ValidateCustom(fl validator.FieldLevel) bool {
	return true
}

func init() {
	UNICRM_VALIDATE = validator.New()
	UNICRM_VALIDATE.RegisterValidation("custom", ValidateCustom)
}
