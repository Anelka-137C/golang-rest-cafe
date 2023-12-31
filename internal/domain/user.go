package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string `json:"name" bson:"name" binding:"required,min=6"`
	Email    string `json:"email" bson:"email" binding:"required,validateEmail,email"`
	Role     string `json:"role" bson:"role" binding:"required,validateRole"`
	Password string `json:"password" bson:"password" binding:"required,min=7,max=13"`
	Active   bool   `default:"true" json:"active" bson:"active"`
}

type UserResponse struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" bson:"name" binding:"required,min=6"`
	Email    string             `json:"email" bson:"email" binding:"required,validateEmail,email"`
	Role     string             `json:"role" bson:"role" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required,min=7,max=10"`
	Active   bool               `default:"true" json:"active" bson:"active"`
}

type Login struct {
	Email    string `json:"email" bson:"email" binding:"required,validateIfExistEmail,email"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type LoginResponse struct {
	Jwt string
}
