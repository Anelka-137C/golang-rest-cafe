package main

import (
	"github.com/Anelka-137C/cafe-app/db"
	"github.com/Anelka-137C/cafe-app/middlewares"

	"github.com/Anelka-137C/cafe-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db.ConnectDB()
	e := echo.New()
	e.Use(middlewares.ValidateEmail)
	e.Use()
	e.POST("/users", routes.CreateUser)
	e.GET("/users", routes.GetAllUsers)
	e.GET("users/:email", routes.GetUser)
	e.DELETE("users/:email", routes.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
