package util

import (
	"errors"
	"os"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	secretKey = "SECRET_KEY"
	adminRole = "ADMIN_ROLE"
)

func BuildResponse(status int, object interface{}, message string, c *gin.Context, entity string) {

	switch object.(type) {
	case domain.LoginResponse:
		c.JSON(status, gin.H{
			"token": object,
		})
		return
	default:
		if object != nil {
			c.JSON(status, gin.H{
				"msg":  message,
				entity: object,
			})
		} else {
			c.JSON(status, gin.H{
				"msg": message,
			})
		}
		return
	}

}

func BuildBadResponse(status int, err []domain.ErrorMsg, c *gin.Context) {
	c.JSON(status, gin.H{"error": err})
}

func ExtractClaimsFromJwt(tokenString string, c *gin.Context) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token := c.GetHeader("token")
	err := godotenv.Load()
	if err != nil {
		return claims, errors.New("Error to load .env")

	}
	sampleSecretKey := os.Getenv(secretKey)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(sampleSecretKey), nil
	})

	if err != nil {
		return claims, errors.New("You must send this field")

	}

	return claims, nil
}
