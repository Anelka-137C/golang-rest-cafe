package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	IdUser      string          `json:"id_user" bson:"id_user" binding:"required"`
	ProductList []ProductInCard `json:"product_list" bson:"product_list" binding:"required,ValidateIsEmptyProducts,ValidateProducts"`
}

type ProductInCard struct {
	IdProduct string `json:"id_product" bson:"id_product" binding:"required"`
	Quantitie int    `json:"quantite" bson:"quantite" binding:"required"`
}

type CardResponse struct {
	ID          primitive.ObjectID `bson:"_id"`
	IdUser      string             `json:"description" bson:"description"`
	ProductList []ProductInCard    `json:"product_list" bson:"product_list"`
}
