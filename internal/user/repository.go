package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	dataBase        = "GoCafe"
	userCollection  = "users"
	rolesCollection = "roles"
	secretKey       = "SECRET_KEY"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateUser(c *gin.Context) (domain.User, []domain.ErrorMsg)
	GetUser(c *gin.Context) (domain.UserResponse, []domain.ErrorMsg)
	GetUserByEmail(c *gin.Context) (domain.UserResponse, []domain.ErrorMsg)
	GetUsersByName(c *gin.Context) ([]domain.UserResponse, []domain.ErrorMsg)
	DeleteUser(c *gin.Context) []domain.ErrorMsg
	UpdateUser(c *gin.Context) []domain.ErrorMsg
	Login(c *gin.Context) (domain.LoginResponse, []domain.ErrorMsg)
	ValidateEmail(email string) bool
	ValidateRole(role string) bool
}

func NewRepository(db *mongo.Client) Repository {

	return &repository{
		db: db,
	}
}

// CreateUser implements Respository.
func (r *repository) CreateUser(c *gin.Context) (domain.User, []domain.ErrorMsg) {

	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	newUser := domain.User{}

	if err := c.ShouldBindJSON(&newUser); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			return newUser, out
		}
	} else {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		newUser.Password = string(hashedPassword)
		userColl.InsertOne(context.TODO(), newUser)
	}

	return newUser, nil
}

// GetUser implements Repository.
func (r *repository) GetUser(c *gin.Context) (domain.UserResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	user := domain.UserResponse{}

	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)
	if user.ID.IsZero() {
		return user, helpers.GenerateOneError("id", "The user is not in data base")
	}

	return user, nil
}

func (r *repository) GetUsersByName(c *gin.Context) ([]domain.UserResponse, []domain.ErrorMsg) {

	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	user := domain.UserResponse{}
	userList := []domain.UserResponse{}
	userName := c.Query("name")
	regularExpr := fmt.Sprintf("%s.*", userName)
	filter := bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: regularExpr}}}}

	cursor, err := userColl.Find(context.TODO(), filter)
	if err != nil {
		return nil, helpers.GenerateOneError("name", "There was an error during the search")
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&user)
		userList = append(userList, user)
	}

	if len(userList) == 0 {
		return nil, helpers.GenerateOneError("name", "There is no users with this name: "+userName)
	}

	return userList, nil
}

func (r *repository) GetUserByEmail(c *gin.Context) (domain.UserResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)

	user := domain.UserResponse{}
	email := c.Query("email")
	if email == "" {
		return user, helpers.GenerateOneError("email", "The param is empty")
	}

	filter := bson.D{{Key: "email", Value: email}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)
	if user.ID.IsZero() {
		return user, helpers.GenerateOneError("email", "The user is not in data base")
	}

	return user, nil
}

func (r *repository) DeleteUser(c *gin.Context) []domain.ErrorMsg {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	id := c.Param("_id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err = userColl.DeleteOne(context.TODO(), filter)
	if err != nil {
		return helpers.GenerateOneError("id", "Error at the moment to delete")
	}

	return nil
}

func (r *repository) UpdateUser(c *gin.Context) []domain.ErrorMsg {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	auxUser := domain.UserResponse{}

	filter := bson.D{{Key: "_id", Value: objectId}}
	userColl.FindOne(context.TODO(), filter).Decode(&auxUser)
	if auxUser.ID.IsZero() {
		return helpers.GenerateOneError("id", "The user is not in data base")
	}

	user := domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		return helpers.GenerateMultipleErrorMsg(err)
	}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "role", Value: user.Role},
		{Key: "password", Value: user.Password},
		{Key: "active", Value: user.Active}}}}

	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return helpers.GenerateOneError("id", "Error at the moment to update")
	}
	return nil
}

func (r *repository) ValidateEmail(email string) bool {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	user := domain.UserResponse{}
	filter := bson.D{{Key: "email", Value: email}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)
	return user.ID.IsZero()
}

func (r *repository) ValidateRole(role string) bool {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(rolesCollection)
	user := domain.UserResponse{}
	filter := bson.D{{Key: "role", Value: role}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)
	return !user.ID.IsZero()
}

func (r *repository) Login(c *gin.Context) (domain.LoginResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	userColl := dataBase.Collection(userCollection)
	user := domain.UserResponse{}
	loginResponse := domain.LoginResponse{}
	login := domain.Login{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sampleSecretKey := os.Getenv(secretKey)
	if err = c.ShouldBindJSON(&login); err != nil {
		return loginResponse, helpers.GenerateMultipleErrorMsg(err)
	}
	filter := bson.D{{Key: "email", Value: login.Email}}
	userColl.FindOne(context.TODO(), filter).Decode(&user)

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return loginResponse, helpers.GenerateOneError("password", "The password is not correct")
	}

	if !user.Active {
		return loginResponse, helpers.GenerateOneError("id", "User is not activate, you must reactivate the account")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Local().UnixMilli()
	claims["authorized"] = true
	claims["user"] = user.ID
	claims["role"] = user.Role

	tokenString, _ := token.SignedString([]byte(sampleSecretKey))
	loginResponse.Jwt = tokenString
	return loginResponse, nil
}
