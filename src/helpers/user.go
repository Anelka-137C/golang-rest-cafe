package helpers

import "github.com/go-playground/validator/v10"

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "validateEmail":
		return "The email address is registered" + fe.Param()
	}

	return "Unknown error"
}
