package main

import (
	"context"
	"fmt"

	"github.com/Anelka-137C/cafe-app/src/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	eng := gin.Default()

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://nacasas:D960bi8NAzg3hKIN@clusteranelka.vkgfe8i.mongodb.net/GoCafe"))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
