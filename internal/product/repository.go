package product

import (
	"context"
	"errors"
	"fmt"

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
	GetAllProduct(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg)
	GetProductByName(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg)
	DeleteProduct(c *gin.Context) []domain.ErrorMsg
	UpdateProduct(c *gin.Context) []domain.ErrorMsg
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
func (r *repository) GetAllProduct(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	productList := []domain.ProductResponse{}
	role := c.Query("role")
	if role == "" {
		return nil, helpers.GenerateOneError("role", "You must send user's role")
	}

	isActive := struct {
		Active bool `json:"active"`
	}{}
	c.ShouldBindJSON(&isActive)
	filter := bson.D{{Key: "active", Value: isActive.Active}}
	cursor, err := productColl.Find(context.TODO(), filter)
	if err != nil {
		return nil, helpers.GenerateOneError("id", "The id is not a mongo id")
	}

	for cursor.Next(context.TODO()) {
		product := domain.ProductResponse{}
		cursor.Decode(&product)
		if role == "ADMIN_ROLE" {
			productList = append(productList, product)
		} else {
			if product.Active {
				productList = append(productList, product)
			}
		}
	}

	return productList, nil
}

func (r *repository) GetProductByName(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	productList := []domain.ProductResponse{}
	productName := c.Query("name")
	role := c.Query("role")
	if role == "" {
		return nil, helpers.GenerateOneError("role", "You must send user's role")
	}
	regularExpre := fmt.Sprintf("%s.*", productName)
	filter := bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: regularExpre}}}}

	cursor, err := productColl.Find(context.TODO(), filter)

	if err != nil {
		return nil, helpers.GenerateOneError("name", "There was an error during the search")
	}

	for cursor.Next(context.TODO()) {
		product := domain.ProductResponse{}
		cursor.Decode(&product)
		if role == "ADMIN_ROLE" {
			productList = append(productList, product)
		} else {
			if product.Active {
				productList = append(productList, product)
			}
		}
	}

	if len(productList) == 0 {
		return nil, helpers.GenerateOneError("name", "There is no product with the name: "+productName)
	}

	return productList, nil
}

// GetProduct implements Repository.
func (r *repository) GetProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg) {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	product := domain.ProductResponse{}
	id := c.Param("_id")
	role := c.Query("role")
	if role == "" {
		return product, helpers.GenerateOneError("role", "You must send user's role")
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	productColl.FindOne(context.TODO(), filter).Decode(&product)

	if product.ID.IsZero() {
		return product, helpers.GenerateOneError("id", "There is no product with id: "+id)
	}

	if role != "ADMIN_ROLE" && !product.Active {
		return product, helpers.GenerateOneError("id", "The product is not activate")
	}

	return product, nil
}

// UpdateProduct implements Repository.
func (r *repository) UpdateProduct(c *gin.Context) []domain.ErrorMsg {
	dataBase := r.db.Database(dataBase)
	productColl := dataBase.Collection(productCollection)
	id := c.Param("_id")
	role := c.Query("role")

	if role != "ADMIN_ROLE" {
		return helpers.GenerateOneError("role", "the "+role+" role does not have permissions to perform this action.")
	} else if role == "" {
		return helpers.GenerateOneError("role", "You must send user's role")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.GenerateOneError("id", "The id is not a mongo id")
	}
	auxProduct := domain.ProductResponse{}

	filter := bson.D{{Key: "_id", Value: objectId}}
	productColl.FindOne(context.TODO(), filter).Decode(&auxProduct)
	if auxProduct.ID.IsZero() {
		return helpers.GenerateOneError("id", "The product is not in data base")
	}

	product := domain.Product{}
	if err := c.ShouldBindJSON(&product); err != nil {
		return helpers.GenerateMultipleErrorMsg(err)
	}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: product.Name},
		{Key: "description", Value: product.Description},
		{Key: "price", Value: product.Price},
		{Key: "category", Value: product.Category},
		{Key: "active", Value: product.Active}}}}

	_, err = productColl.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return helpers.GenerateOneError("id", "Error at the moment to update")
	}
	return nil
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
