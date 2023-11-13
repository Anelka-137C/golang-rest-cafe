package card

import (
	"context"
	"errors"

	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/Anelka-137C/cafe-app/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dataBase       = "GoCafe"
	CardCollection = "card"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateCard(c *gin.Context) (domain.Card, []domain.ErrorMsg)
	GetCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg)
	DeleteCard(c *gin.Context) []domain.ErrorMsg
	UpdateCard(c *gin.Context) []domain.ErrorMsg
	ValidateArticle(a []domain.Article) bool
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

// CreateCard implements Repository.
func (r *repository) CreateCard(c *gin.Context) (domain.Card, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(CardCollection)
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
		productColl.InsertOne(context.TODO(), newCard)
	}

	return newCard, nil
}

// DeleteCard implements Repository.
func (r *repository) DeleteCard(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// GetCard implements Repository.
func (r *repository) GetCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg) {
	panic("unimplemented")
}

// UpdateCard implements Repository.
func (r *repository) UpdateCard(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

func (r *repository) ValidateArticle(article []domain.Article) bool {

	return true
}
