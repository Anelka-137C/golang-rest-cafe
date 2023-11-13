package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	IdUser      string    `json:"id_user" bson:"id_user" binding:"required"`
	ArticleList []Article `json:"article_list" bson:"article_list" binding:"ValidateArticle"`
}
type Article struct {
	IdProduct string `json:"id_product" bson:"id_product" binding:"required"`
	Quantitie int    `json:"quantite" bson:"quantite" binding:"required"`
}

type CardResponse struct {
	ID          primitive.ObjectID `bson:"_id"`
	IdUser      string             `json:"description" bson:"description"`
	ArticleList []Article          `json:"article_list" bson:"article_list"`
}
