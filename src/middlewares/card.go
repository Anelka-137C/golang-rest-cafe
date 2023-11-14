package middlewares

import (
	"github.com/Anelka-137C/cafe-app/internal/card"
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/go-playground/validator/v10"
)

func ValidateIsEmptyProducts() validator.Func {

	return func(fl validator.FieldLevel) bool {
		article := fl.Field().Interface().([]domain.ProductInCard)
		if len(article) == 0 {
			return false
		} else {
			return true
		}
	}
}

func ValidateProducts(db card.Repository) validator.Func {

	return func(fl validator.FieldLevel) bool {
		article := fl.Field().Interface().([]domain.ProductInCard)
		return db.ValidateArticle(article)
	}
}
