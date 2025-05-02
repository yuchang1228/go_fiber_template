package util

import (
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var (
	Validate *validator.Validate
	trans    ut.Translator
)

func ValidateStruct() {

	zhTw := zh_Hant_TW.New()
	uni := ut.New(zhTw, zhTw)
	trans, _ = uni.GetTranslator("zh_tw")

	Validate = validator.New()

	translations.RegisterDefaultTranslations(Validate, trans)
}

func TranslateErrors(errs validator.ValidationErrors, fieldMap map[string]string) map[string]string {
	translatedErrors := make(map[string]string)
	for _, err := range errs {
		translatedError := err.Translate(trans)
		translatedErrors[err.Field()] = translatedError
	}
	return translatedErrors
}
