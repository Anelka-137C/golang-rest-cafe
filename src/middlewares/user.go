package middlewares

import (
	"net/http"
	"time"

	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	secretKey = "SECRET_KEY"
	adminRole = "ADMIN_ROLE"
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

func ValidateIfExistEmail(db user.Repository) validator.Func {
	return func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		return !db.ValidateEmail(email)
	}
}

func ValidateJwt() gin.HandlerFunc {
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
		if claims["role"] != adminRole {
			errors := helpers.GenerateOneError("token", "User is not ADMIN ROLE")
			util.BuildBadResponse(http.StatusBadRequest, errors, c)
			c.Abort()
			return
		}

	}
}
