package product

import (
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
)

type service struct {
	repository Repository
}

type Service interface {
	CreateProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg)
	GetProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg)
	GetAllProduct(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg)
	GetProductByName(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg)
	DeleteProduct(c *gin.Context) []domain.ErrorMsg
	UpdateProduct(c *gin.Context) []domain.ErrorMsg
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// CreateProduct implements Service.
func (s *service) CreateProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	return s.repository.CreateProduct(c)
}

// DeleteProduct implements Service.
func (s *service) DeleteProduct(c *gin.Context) []domain.ErrorMsg {
	return s.repository.DeleteProduct(c)
}

// GetAllProduct implements Service.
func (s *service) GetAllProduct(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg) {
	return s.repository.GetAllProduct(c)
}

func (s *service) GetProductByName(c *gin.Context) ([]domain.ProductResponse, []domain.ErrorMsg) {
	return s.repository.GetProductByName(c)
}

// GetProduct implements Service.
func (s *service) GetProduct(c *gin.Context) (domain.ProductResponse, []domain.ErrorMsg) {
	return s.repository.GetProduct(c)
}

// UpdateProduct implements Service.
func (s *service) UpdateProduct(c *gin.Context) []domain.ErrorMsg {
	return s.repository.UpdateProduct(c)
}
