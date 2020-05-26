package public

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
	"strings"
)

var (
	Validate *validator.Validate
)

//初始化语言包
func InitValidate() {
	Validate = validator.New()
	if err := Validate.RegisterValidation("sql", ValidateMysqlInjection); err != nil {
		log.Fatalf(" [ERROR] Validate add sql err:%v\n", err)
	}
	log.Printf(" [INFO] Validate add sql\n")
}

func ValidateMysqlInjection(f1 validator.FieldLevel) bool {
	if strings.Index(f1.Field().String(), "=") > -1 {
		return false
	}
	if strings.Index(f1.Field().String(), ",") > -1 {
		return false
	}
	if strings.Index(f1.Field().String(), "-") > -1 {
		return false
	}
	if strings.Index(f1.Field().String(), ")") > -1 {
		return false
	}
	if strings.Index(f1.Field().String(), "(") > -1 {
		return false
	}
	if strings.Index(f1.Field().String(), "or") > -1 {
		return false
	}
	if strings.Index(f1.Field().String(), "*") > -1 {
		return false
	}
	return true
}
