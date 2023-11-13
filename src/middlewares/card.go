package middlewares

import (
	"github.com/Anelka-137C/cafe-app/internal/card"
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/go-playground/validator/v10"
)

func ValidateArticle(db card.Repository) validator.Func {

	return func(fl validator.FieldLevel) bool {
		article := fl.Field().Interface().([]domain.Article)
		return db.ValidateArticle(article)
	}
}
