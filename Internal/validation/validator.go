package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct using validator/v10 tags
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
