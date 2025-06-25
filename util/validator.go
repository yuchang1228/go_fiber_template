package util

import (
	"reflect"

	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewValidator(fieldMap map[string]string) *Validator {
	zhTw := zh_Hant_TW.New()
	uni := ut.New(zhTw, zhTw)
	trans, _ := uni.GetTranslator("zh_tw")

	v := validator.New()

	translations.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if name, ok := fieldMap[fld.Name]; ok {
			return name
		}
		return fld.Name
	})

	return &Validator{
		validate:   v,
		translator: trans,
	}
}

func (v *Validator) ValidateStruct(s interface{}) []string {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	var messages []string
	for _, e := range errs {
		messages = append(messages, e.Translate(v.translator))
	}
	return messages
}
