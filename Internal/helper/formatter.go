package validatorhelper

import (
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
)

// FieldJSONFormatter function to create field name same with json name
func FieldJSONFormatter(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
func RequiredFormatter(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is Required!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
}
func EmailFormatter(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} is not valid email format!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})
}
func EqualPassword(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
		return ut.Add("eqfield", "{0} not same with {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eqfield", fe.Field(), fe.Param())

		return t
	})
}
