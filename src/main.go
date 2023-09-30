package main

import (
	"github.com/Anelka-137C/cafe-app/src/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	rUser := r.Group("/user")
	rUser.GET("/ping", handlers.Pong())

	r.Run() // listen and serve on 0.0.0.0:8080
}
