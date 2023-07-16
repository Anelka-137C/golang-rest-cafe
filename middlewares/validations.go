package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Anelka-137C/cafe-app/db"
	"github.com/Anelka-137C/cafe-app/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type ErrorMessage struct {
	msg string
}

func ValidateEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user models.User
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		json.NewDecoder(c.Request().Body).Decode(&user)
		filter := bson.D{{Key: "email", Value: user.Email}}
		result := db.UserColl.FindOne(ctx, filter)
		if result != nil {
			return c.String(http.StatusBadRequest, "Usuario con el email ingresado ya existe")
		}

		return nil
	}
}
