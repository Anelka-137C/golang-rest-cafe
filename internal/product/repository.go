package product

import (
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
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
	GetProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg)
	GetAllProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg)
	DeleteProduct(c *gin.Context) []domain.ErrorMsg
	UpdateProduct(c *gin.Context) []domain.ErrorMsg
	ValidateName(c *gin.Context) bool
	ValidateCategory(c *gin.Context) bool
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

// CreateProduct implements Repository.
func (r *repository) CreateProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	// dataBase := r.db.Database(dataBase)
	// userColl := dataBase.Collection(productCollection)
	newProduct := domain.Product{
		Name:        "Hola",
		Description: "Hola",
		Price:       23,
		Category:    "Hola",
	}

	return newProduct, nil
}

// DeleteProduct implements Repository.
func (r *repository) DeleteProduct(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// GetAllProduct implements Repository.
func (r *repository) GetAllProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	panic("unimplemented")
}

// GetProduct implements Repository.
func (r *repository) GetProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	panic("unimplemented")
}

// UpdateProduct implements Repository.
func (r *repository) UpdateProduct(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// ValidateCategory implements Repository.
func (r *repository) ValidateCategory(c *gin.Context) bool {
	panic("unimplemented")
}

// ValidateName implements Repository.
func (r *repository) ValidateName(c *gin.Context) bool {
	panic("unimplemented")
}
