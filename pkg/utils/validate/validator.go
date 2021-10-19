package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"home/pkg/code"
	"home/pkg/code/e"
	"log"
	"strings"
)

var (
	v     *validator.Validate
	trans ut.Translator
)

func init() {
	v = validator.New()
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	err := translations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		log.Println(err.Error())
	}
}

// Struct validate params struct
func Struct(params interface{}) error {
	if err := v.Struct(params); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range validationErrors {
				return e.NewError(code.ParamsIsInvalid, err.Translate(trans))
			}
		}
	}

	return nil
}

// Variable Single parameter
func Variable(fieldName string, val interface{}, tag string) error {
	if err := v.Var(val, strings.ReplaceAll(tag, " ", "")); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range validationErrors {
				return e.NewError(code.ParamsIsInvalid, fieldName+err.Translate(trans))
			}
		}
	}

	return nil
}
