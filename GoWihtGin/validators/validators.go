package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

func ValidateCoolTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Cool")
}
