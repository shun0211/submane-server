package utils

import (
	"reflect"

	ja "github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	ja_translations "gopkg.in/go-playground/validator.v9/translations/ja"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	ja := ja.New()
	uni = ut.New(ja, ja)

	trans, _ := uni.GetTranslator("ja")
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		fieldName := fld.Tag.Get("jaFieldName")
		if fieldName == "-" {
			return ""
		}

		return fieldName
	})

	ja_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		validator: validate,
	}
}

func(v *Validator) Validate(i interface{}) (err error) {
	return v.validator.Struct(i)
}

func GetErrorMessages(err error) (messages []string) {
	trans, _ := uni.GetTranslator("ja")
	for _, m := range err.(validator.ValidationErrors).Translate(trans) {
		messages = append(messages, m)
	}
	return messages
}
