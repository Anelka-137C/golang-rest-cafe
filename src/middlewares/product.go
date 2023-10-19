package middlewares

import (
	"github.com/Anelka-137C/cafe-app/internal/product"
	"github.com/go-playground/validator/v10"
)

func ValidateCategory(db product.Repository) validator.Func {

	return func(fl validator.FieldLevel) bool {
		category := fl.Field().String()
		return db.ValidateCategory(category)
	}
}
