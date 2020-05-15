package Validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/open-kingfisher/king-utils/common/log"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("requiredValidate", requiredValidate); err != nil {
			log.Errorf("RegisterValidation error:%v", err)
		}
	}
}

var requiredValidate validator.Func = func(fl validator.FieldLevel) bool {
	if data, ok := fl.Field().Interface().([]string); ok && len(data) > 0 {
		return true
	}
	if data, ok := fl.Field().Interface().(string); ok && data != "" {
		return true
	}
	return false
}
