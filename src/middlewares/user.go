package middlewares

import (
	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/go-playground/validator/v10"
)

func ValidateEmail(db user.Repository) validator.Func {

	return func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		return db.ValidateEmail(email)
	}
}

func ValidateRole(db user.Repository) validator.Func {
	return func(fl validator.FieldLevel) bool {
		role := fl.Field().String()
		return db.ValidateRole(role)
	}
}
