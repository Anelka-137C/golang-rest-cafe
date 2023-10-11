package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateUser(c *gin.Context) (domain.User, error)
	GetUser(c *gin.Context) (domain.User, error)
	DeleteUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	ValidateEmail(email string) bool
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

	if err := c.ShouldBindJSON(&newUser); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		} else {
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 9)
			newUser.Password = string(hashedPassword)
			userColl.InsertOne(context.TODO(), newUser)
		}
	}

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

func (r *repository) ValidateEmail(email string) bool {
	dataBase := r.db.Database("GoCafe")
	userColl := dataBase.Collection("users")
	user := domain.UserResponse{}
	filter := bson.D{{Key: "email", Value: email}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)

	return user.ID.IsZero()
}
