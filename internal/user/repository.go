package user

import (
	"context"
	"fmt"
	"log"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateUser(c *gin.Context) domain.User
	GetUser(c *gin.Context) domain.User
	DeleteUser(c *gin.Context)
}

func NewRepository(db *mongo.Client) Repository {

	return &repository{
		db: db,
	}
}

// CreateUser implements Respository.
func (r *repository) CreateUser(c *gin.Context) domain.User {

	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	newUser := domain.User{}
	err := c.Bind(&newUser)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(newUser)
	userColl.InsertOne(context.TODO(), newUser)

	return newUser
}

// GetUser implements Repository.
func (r *repository) GetUser(c *gin.Context) domain.User {
	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	user := domain.User{}
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)
	return user
}

func (r *repository) DeleteUser(c *gin.Context) {
	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	userColl.DeleteOne(context.TODO(), filter)

}
