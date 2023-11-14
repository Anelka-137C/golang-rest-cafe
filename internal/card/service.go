package card

import (
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
)

type service struct {
	repository Repository
}

type Service interface {
	CreateCard(c *gin.Context) (domain.Card, []domain.ErrorMsg)
	GetCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg)
	DeleteCard(c *gin.Context) []domain.ErrorMsg
	AddToCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// CreateCard implements Service.
func (s *service) CreateCard(c *gin.Context) (domain.Card, []domain.ErrorMsg) {
	return s.repository.CreateCard(c)
}

// DeleteCard implements Service.
func (s *service) DeleteCard(c *gin.Context) []domain.ErrorMsg {
	panic("unimplemented")
}

// GetCard implements Service.
func (s *service) GetCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg) {
	return s.repository.GetCard(c)
}

// UpdateCard implements Service.
func (s *service) AddToCard(c *gin.Context) (domain.CardResponse, []domain.ErrorMsg) {
	return s.repository.AddToCard(c)
}
