package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Anelka-137C/cafe-app/db"
	"github.com/Anelka-137C/cafe-app/middlewares"
	"github.com/Anelka-137C/cafe-app/models"
)

func CreateUser(c echo.Context) error {

	user := middlewares.User
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	json.NewDecoder(c.Request().Body).Decode(&user)
	fmt.Println("usuario en el metodo de insercion: ", user)
	user.Active = true
	result, insertErr := db.UserColl.InsertOne(ctx, user)
	if insertErr != nil {
		fmt.Println("Error al insertar:", insertErr)
	} else {
		fmt.Println("Tipo de insercion: ", reflect.TypeOf(result))
		fmt.Println("resultado de insercion de insercion: ", result)
	}
	return json.NewEncoder(c.Response()).Encode(&user)
}

func GetAllUsers(c echo.Context) error {

	var users []models.User
	var cursor *mongo.Cursor
	ctx, ctxErr := context.WithTimeout(context.Background(), 15*time.Second)
	if ctxErr != nil {
		panic("Error al cargar contexto")
	}
	cursor.All(ctx, &users)

	return json.NewEncoder(c.Response()).Encode(&users)

}
