package card

import (
	"context"
	"errors"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dataBase          = "GoCafe"
	CardCollection    = "card"
	productCollection = "product"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateCard(c *gin.Context) (domain.Card, []domain.ErrorMsg)
	GetCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg)
	AddToCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg)
	DeleteCard(c *gin.Context) []domain.ErrorMsg
	ValidateArticle(articles []domain.ProductInCard) bool
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

// CreateCard implements Repository.
func (r *repository) CreateCard(c *gin.Context) (domain.Card, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	cardColl := dataBase.Collection(CardCollection)
	newCard := domain.Card{}

	if err := c.ShouldBindJSON(&newCard); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			return newCard, out
		}
	} else {
		cardColl.InsertOne(context.TODO(), newCard)
	}

	return newCard, nil
}

// DeleteCard implements Repository.
func (r *repository) DeleteCard(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// GetCard implements Repository.
func (r *repository) GetCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	cardColl := dataBase.Collection(CardCollection)
	cardResponse := domain.CardResponse{}

	idUser := c.Param("_id")
	filter := bson.D{{Key: "id_user", Value: idUser}}
	cardColl.FindOne(context.TODO(), filter).Decode(&cardResponse)
	return cardResponse, nil
}

func (r *repository) ValidateArticle(articles []domain.ProductInCard) bool {
	dataBase := r.db.Database(dataBase)
	cardColl := dataBase.Collection(productCollection)
	productToValidate := domain.ProductResponse{}

	for _, article := range articles {
		objectId, _ := primitive.ObjectIDFromHex(article.IdProduct)

		filter := bson.D{{Key: "_id", Value: objectId}}
		cardColl.FindOne(context.TODO(), filter).Decode(&productToValidate)
		if productToValidate.ID.IsZero() {
			return false
		}
	}
	return true
}

func (r *repository) AddToCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	cardColl := dataBase.Collection(CardCollection)
	cardResponse := domain.CardResponse{}
	newCardProduct := domain.ProductInCard{}

	if err := c.ShouldBindJSON(&newCardProduct); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			return cardResponse, out
		}
	}

	idUser := c.Param("_id")
	filter := bson.D{{Key: "id_user", Value: idUser}}
	cardColl.FindOne(context.TODO(), filter).Decode(&cardResponse)

	productList := cardResponse.ProductList

	contains, index := verifyContainsProduct(newCardProduct.IdProduct, productList)
	if contains {

		cardProduct := productList[index]
		cardProduct.Quantitie = cardProduct.Quantitie + newCardProduct.Quantitie
		productList[index] = cardProduct
		cardResponse.ProductList = productList
	} else {
		productList = append(productList, newCardProduct)
		cardResponse.ProductList = productList
	}

	updatedList := bson.D{{Key: "product_list", Value: productList}}
	update := bson.D{{Key: "$set", Value: updatedList}}
	cardColl.UpdateOne(context.TODO(), filter, update)

	return cardResponse, nil

}

func verifyContainsProduct(id string, productList []domain.ProductInCard) (bool, int) {

	for index, product := range productList {
		if product.IdProduct == id {
			return true, index
		}
	}

	return false, 0
}
