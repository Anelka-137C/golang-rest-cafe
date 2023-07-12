package routes

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/Anelka-137C/cafe-app/models"

	"net/http"
)

func CreateUser(c echo.Context) error {

	var user models.User
	json.NewDecoder(c.Request().Body).Decode(&user)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(&user)
}
