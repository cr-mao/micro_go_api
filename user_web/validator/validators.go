package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)
//验证手机
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	return ok
}
