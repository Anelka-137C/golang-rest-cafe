package product

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
	dataBase             = "GoCafe"
	productCollection    = "product"
	categoriesCollection = "category"
)

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg)
	GetProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg)
	GetAllProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg)
	DeleteProduct(c *gin.Context) []domain.ErrorMsg
	UpdateProduct(c *gin.Context) []domain.ErrorMsg
	ValidateName(name string) bool
	ValidateCategory(category string) bool
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

// CreateProduct implements Repository.
func (r *repository) CreateProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	newProduct := domain.Product{}

	if err := c.ShouldBindJSON(&newProduct); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]domain.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = domain.ErrorMsg{Field: fe.Field(), Message: helpers.GetErrorMsg(fe)}
			}
			return newProduct, out
		}
	} else {

		productColl.InsertOne(context.TODO(), newProduct)
	}

	return newProduct, nil
}

// DeleteProduct implements Repository.
func (r *repository) DeleteProduct(c *gin.Context) []domain.ErrorMsg {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	id := c.Param("_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err = productColl.DeleteOne(context.TODO(), filter)
	if err != nil {
		return helpers.GenerateOneError("id", "Error at the moment to delete")
	}

	return nil
}

// GetAllProduct implements Repository.
func (r *repository) GetAllProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg) {
	panic("unimplemented")
}

// GetProduct implements Repository.
func (r *repository) GetProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	id := c.Param("_id")
	product := domain.ProductResponse{}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	productColl.FindOne(context.TODO(), filter).Decode(&product)

	return product, nil
}

// UpdateProduct implements Repository.
func (r *repository) UpdateProduct(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// ValidateCategory implements Repository.
func (r *repository) ValidateCategory(category string) bool {
	dataBase := r.db.Database(dataBase)
	categoryColl := dataBase.Collection(categoriesCollection)
	categoryToValidate := domain.Category{}
	filter := bson.D{{Key: "name", Value: category}}
	categoryColl.FindOne(context.TODO(), filter).Decode(&categoryToValidate)
	return !categoryToValidate.ID.IsZero()
}

// ValidateName implements Repository.
func (r *repository) ValidateName(name string) bool {
	panic("unimplemented")
}
