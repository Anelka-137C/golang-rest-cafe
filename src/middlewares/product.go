package middlewares

import (
	"net/http"
	"time"

	"github.com/Anelka-137C/cafe-app/internal/product"
	"github.com/gin-gonic/gin"

	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/Anelka-137C/cafe-app/src/util"

	"github.com/go-playground/validator/v10"
)

func ValidateCategory(db product.Repository) validator.Func {

	return func(fl validator.FieldLevel) bool {
		category := fl.Field().String()
		return db.ValidateCategory(category)
	}
}

func ValidateJwtForUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		date := time.Now()
		token := c.GetHeader("token")

		claims, err := util.ExtractClaimsFromJwt(token, c)

		if err != nil {
			errors := helpers.GenerateOneError("token", err.Error())
			util.BuildBadResponse(http.StatusBadRequest, errors, c)
			c.Abort()
			return
		}
		if !claims.VerifyExpiresAt(date.UnixMilli(), true) {
			errors := helpers.GenerateOneError("token", "Session expired")
			util.BuildBadResponse(http.StatusBadRequest, errors, c)
			c.Abort()
			return
		}

	}
}
