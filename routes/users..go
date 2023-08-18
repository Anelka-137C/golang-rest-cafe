package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

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
	return c.JSON(http.StatusOK, &user)
}

func GetAllUsers(c echo.Context) error {

	var users []models.User

	filter := bson.D{{"active", true}}
	cursor, err := db.UserColl.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, &users)
}

func GetUser(c echo.Context) error {
	var user models.User
	email := c.Param("email")
	filter := bson.D{{"email", email}}
	db.UserColl.FindOne(context.TODO(), filter).Decode(&user)
	return c.JSON(http.StatusOK, &user)
}

func DeleteUser(c echo.Context) error {
	msg := models.ResponseDelete{
		Msg: "Registro eliminado con exito",
	}
	email := c.Param("email")
	filter := bson.D{{"email", email}}
	db.UserColl.DeleteOne(context.TODO(), filter)

	return c.JSON(http.StatusOK, &msg)
}
