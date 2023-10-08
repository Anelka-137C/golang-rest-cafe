package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Anelka-137C/cafe-app/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoUrl = "MONGODB_URI"

func main() {

	eng := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoUri := os.Getenv(mongoUrl)
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error to connect DB")
	}

	router := routes.NewRouter(eng, db)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}
