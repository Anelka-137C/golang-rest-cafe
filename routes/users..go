package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Anelka-137C/cafe-app/db"
	"github.com/Anelka-137C/cafe-app/models"
)

func CreateUser(c echo.Context) error {

	var user models.User
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	json.NewDecoder(c.Request().Body).Decode(&user)
	result, insertErr := db.UserColl.InsertOne(ctx, user)
	if insertErr != nil {
		fmt.Println("Error al insertar:", insertErr)
	} else {
		fmt.Println("Tipo de insercion: ", reflect.TypeOf(result))
		fmt.Println("resultado de insercion de insercion: ", result)
	}
	return json.NewEncoder(c.Response()).Encode(&user)
}
