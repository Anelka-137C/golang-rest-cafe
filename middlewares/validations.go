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

var User models.User

func ValidateEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var userFromDB models.User
		userInfo := c.Request().Body
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

		json.NewDecoder(userInfo).Decode(&User)
		filter := bson.D{{Key: "email", Value: User.Email}}
		db.UserColl.FindOne(ctx, filter).Decode(&userFromDB)

		if userFromDB != (models.User{}) {
			return c.String(http.StatusBadRequest, "Usuario con el email ingresado ya existe")
		}

		return next(c)
	}
}
