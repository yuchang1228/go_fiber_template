package validator

import (
	"app/pkg/i18n"
	"reflect"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

var uni *ut.UniversalTranslator

func initValidator() {
	en := en.New()
	zhTw := zh_Hant_TW.New()

	uni = ut.New(zhTw, zhTw, en)
}

// 建立 Validator 實例
func NewValidator(lang string, fieldMap ...map[string]string) *Validator {
	initValidator()

	trans, _ := uni.GetTranslator(lang)

	validate := validator.New()

	registerTranslations(validate, lang, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if len(fieldMap) > 0 {
			if name, ok := fieldMap[0][fld.Name]; ok {
				return name
			}
		}

		if name, err := i18n.Localize(lang, fld.Name); err == nil {
			return name
		}

		return fld.Name
	})

	return &Validator{
		validate:   validate,
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

func registerTranslations(validate *validator.Validate, lang string, trans ut.Translator) {
	switch lang {
	case "en":
		en_translations.RegisterDefaultTranslations(validate, trans)
	case "zh_tw":
		zh_tw_translations.RegisterDefaultTranslations(validate, trans)
	default:
		zh_tw_translations.RegisterDefaultTranslations(validate, trans)
	}
}
