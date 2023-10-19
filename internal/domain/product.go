package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Name        string  `json:"name" bson:"name" binding:"required"`
	Description string  `json:"description" bson:"description" binding:"required"`
	Price       float32 `json:"price" bson:"price" binding:"required"`
	Category    string  `json:"category" bson:"category" binding:"required,ValidateCategory"`
}

type ProductResponse struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	Price       float32            `json:"price" bson:"price" binding:"required"`
	Category    string             `json:"category" bson:"category" binding:"category"`
}

type Category struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `json:"name" bson:"name"`
}
