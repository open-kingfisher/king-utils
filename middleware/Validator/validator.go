package Validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/open-kingfisher/king-utils/common/log"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("requiredValidate", requiredValidate); err != nil {
			log.Errorf("RegisterValidation error:%v", err)
		}
	}
}

func requiredValidate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if data, ok := field.Interface().([]string); ok && len(data) > 0 {
		return true
	}
	if data, ok := field.Interface().(string); ok && data != "" {
		return true
	}
	return false
}
