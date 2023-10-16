package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
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
		claims := jwt.MapClaims{}
		date := time.Now()
		token := c.GetHeader("token")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		sampleSecretKey := os.Getenv(secretKey)
		_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(sampleSecretKey), nil
		})
		if err != nil {
			log.Fatal(err)
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
