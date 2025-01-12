package validatorhelper

import (
	en_EN "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// MessageError DataStructure MessageError with field and message
type MessageError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type CustomValidator struct {
	trans     ut.Translator
	validator *validator.Validate
}

func NewValidator(validatorVal *validator.Validate) CustomValidator {
	return CustomValidator{
		validator: validatorVal,
	}
}

func (v *CustomValidator) Validate(i interface{}) []MessageError {
	var messages []MessageError

	err := v.validator.Struct(i)
	if err != nil {
		return v.errorMessageTranslator(err, v.trans)

	}
	return messages
}

// ErrorMessageTranslator function to customize the error message
func (v *CustomValidator) errorMessageTranslator(err error, translated ut.Translator) []MessageError {
	var messages []MessageError
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		var messageStruct MessageError
		messageStruct.Field = e.Field()
		messageStruct.Message = e.Translate(translated)
		messages = append(messages, messageStruct)
		// can translated each error one at a time.

	}
	return messages
}

// RequiredErrorMessage function to replace error message for  "required" type
func (v *CustomValidator) RequiredErrorMessage() {
	// NOTE: ommitting allot of error checking for brevity

	english := en_EN.New()
	uni = ut.New(english, english)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	en.RegisterDefaultTranslations(v.validator, trans)
	FieldJSONFormatter(v.validator)
	RequiredFormatter(v.validator, trans)
	EmailFormatter(v.validator, trans)
	EqualPassword(v.validator, trans)
	v.trans = trans
}
