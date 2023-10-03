package user

import (
	"context"
	"fmt"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateUser(c *gin.Context) domain.User
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
