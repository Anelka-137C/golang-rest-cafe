package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

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
		resUserRepeated := ErrorMessage{
			msg: "Usuario ya esta registrado",
		}
		json.NewDecoder(c.Request().Body).Decode(&user)

		collection := db.DB.Database("GoCafeDB").Collection("users")

		err := collection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}})

		if err != nil {
			return c.JSON(http.StatusConflict, resUserRepeated)
		}

		return nil
	}
}
