package common

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var ValiObj *validator.Validate
var ValiTrans ut.Translator

//https://godoc.org/gopkg.in/go-playground/validator.v9
func init() {
	ValiObj = validator.New()
	english := en.New()
	chinese := zh.New()
	uni := ut.New(chinese, english)
	ValiTrans, _ = uni.GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(ValiObj, ValiTrans)
}

//检测是否合法
func Validator(params interface{}) error {
	err := ValiObj.Struct(params)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(ValiTrans))
		}
	}

	return nil
}
