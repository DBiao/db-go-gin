package validata

import (
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 表单验证
// https://github.com/go-playground/validator
var (
	Validate      *validator.Validate
	uni           *ut.UniversalTranslator
	validateTrans ut.Translator
)

func InitValid() error {
	zh2 := zh.New()
	uni = ut.New(zh2, zh2)

	validateTrans, _ = uni.GetTranslator("zh")
	Validate = validator.New()

	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("comment")
	})

	if err := zh_translations.RegisterDefaultTranslations(Validate, validateTrans); err != nil {
		return err
	}

	return nil
}

func ValidRequest(validErr interface{}) []string {
	var errs []string
	if validErr == nil {
		return errs
	}

	validErrs := validErr.(validator.ValidationErrors)
	for _, e := range validErrs.Translate(validateTrans) {
		if len(e) == 0 {
			continue
		}
		errs = append(errs, e)
	}

	return errs
}
