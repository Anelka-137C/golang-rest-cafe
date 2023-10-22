package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Anelka-137C/cafe-app/internal/product"
	"github.com/gin-gonic/gin"

	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/dgrijalva/jwt-go"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func ValidateCategory(db product.Repository) validator.Func {

	return func(fl validator.FieldLevel) bool {
		category := fl.Field().String()
		return db.ValidateCategory(category)
	}
}

func ValidateJwtForUsers() gin.HandlerFunc {
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
			errors := helpers.GenerateOneError("token", "You must send this field")
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
