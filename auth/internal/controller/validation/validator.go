package validation

import (
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("password", passValidate)
}

func passValidate(fl validator.FieldLevel) bool {
	// там не может быть не строка
	pass := fl.Field().String()

	if utf8.RuneCountInString(pass) < 8 {
		return false
	}

	var (
		upperCase bool
		lowerCase bool
		digit     bool
	)

	for _, uni := range pass {
		switch {
		case unicode.IsUpper(uni):
			upperCase = true
		case unicode.IsLower(uni):
			lowerCase = true
		case unicode.IsDigit(uni):
			digit = true
		}
	}

	if upperCase && lowerCase && digit {
		return true
	} else {
		return false
	}
}
