package helpers

import (
	"errors"
	"fmt"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/go-playground/validator/v10"
)

func GenerateMultipleErrorMsg(err error) []domain.ErrorMsg {
	var ve validator.ValidationErrors
	errors.As(err, &ve)
	out := make([]domain.ErrorMsg, len(ve))
	for i, fe := range ve {
		out[i] = domain.ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
	}

	return out
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "validateEmail":
		return "The email " + fmt.Sprintf("%s", fe.Value()) + " is registered"
	case "validateRole":
		return "The Role " + fmt.Sprintf("%s", fe.Value()) + " is not defined in data base"
	case "email":
		return "The format of " + fmt.Sprintf("%s", fe.Value()) + "does not correspond to an email address"
	case "min":
		return "Should be more than " + fe.Param()
	case "max":
		return "Should be greater than " + fe.Param()
	case "validateIfExistEmail":
		return "The email " + fmt.Sprintf("%s", fe.Value()) + " is not registered"
	case "ValidateCategory":
		return "The category " + fmt.Sprintf("%s", fe.Value()) + " is not registered"
	}

	return "Unknown error"
}

func GenerateOneError(field, message string) []domain.ErrorMsg {
	errMessages := make([]domain.ErrorMsg, 1)
	errMessages[0].Field = field
	errMessages[0].Message = message

	return errMessages
}
