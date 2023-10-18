package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float32 `json:"price" bson:"price"`
	Category    string  `json:"category" bson:"category"`
}

type ProductResponse struct {
	ID          primitive.ObjectID `bson:"_id"`
	name        string             `json:"name" bson:"name" binding:"required"`
	description string             `json:"description" bson:"description" binding:"required"`
	price       float32            `json:"price" bson:"price" binding:"required"`
	category    string             `json:"category" bson:"category" binding:"category"`
}
