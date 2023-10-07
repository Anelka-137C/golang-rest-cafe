package user

import (
	"context"
	"fmt"

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
	CreateUser(c *gin.Context) (domain.User, error)
	GetUser(c *gin.Context) (domain.User, error)
	DeleteUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
}

func NewRepository(db *mongo.Client) Repository {

	return &repository{
		db: db,
	}
}

// CreateUser implements Respository.
func (r *repository) CreateUser(c *gin.Context) (domain.User, error) {

	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	newUser := domain.User{}

	err := c.ShouldBindJSON(&newUser)

	if err != nil {
		return newUser, err
	}

	fmt.Println(newUser)
	userColl.InsertOne(context.TODO(), newUser)

	return newUser, nil
}

// GetUser implements Repository.
func (r *repository) GetUser(c *gin.Context) (domain.User, error) {
	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	user := domain.User{}
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)
	return user, nil
}

func (r *repository) DeleteUser(c *gin.Context) error {
	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	userColl.DeleteOne(context.TODO(), filter)
	return nil
}

func (r *repository) UpdateUser(c *gin.Context) error {
	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user := domain.User{}
	c.Bind(&user)
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "role", Value: user.Role},
		{Key: "password", Value: user.Password},
		{Key: "active", Value: user.Active}}}}

	filter := bson.D{{Key: "_id", Value: objectId}}

	userColl.UpdateOne(context.TODO(), filter, update)
	return nil
}
