package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

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
	}

	return "Unknown error"
}
