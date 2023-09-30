package handlers

import (
	"context"

	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Pong() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func CreateUser(c *gin.Context) func(c *gin.Context) {

	db, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://nacasas:D960bi8NAzg3hKIN@clusteranelka.vkgfe8i.mongodb.net/GoCafe"))
	repo := user.NewRepository(db)
	service := user.NewService(repo)

	return func(c *gin.Context) {
		service.CreateUser(c)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}

}
