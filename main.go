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
	e.GET("/users", routes.CreateUser, middlewares.ValidateEmail)
	e.Logger.Fatal(e.Start(":1323"))
}
