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
	GetProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg)
	GetAllProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg)
	DeleteProduct(c *gin.Context) []domain.ErrorMsg
	UpdateProduct(c *gin.Context) []domain.ErrorMsg
	ValidateName(c *gin.Context) bool
	ValidateCategory(c *gin.Context) bool
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
	panic("unimplemented")
}

// GetAllProduct implements Service.
func (s *service) GetAllProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	panic("unimplemented")
}

// GetProduct implements Service.
func (s *service) GetProduct(c *gin.Context) (domain.Product, []domain.ErrorMsg) {
	panic("unimplemented")
}

// UpdateProduct implements Service.
func (s *service) UpdateProduct(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// ValidateCategory implements Service.
func (s *service) ValidateCategory(c *gin.Context) bool {
	panic("unimplemented")
}

// ValidateName implements Service.
func (s *service) ValidateName(c *gin.Context) bool {
	panic("unimplemented")
}
