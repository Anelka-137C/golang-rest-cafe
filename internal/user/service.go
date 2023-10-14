package user

import (
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
)

type service struct {
	repository Repository
}

// CreateUser implements Service.

type Service interface {
	CreateUser(c *gin.Context) (domain.User, []domain.ErrorMsg)
	GetUser(c *gin.Context) (domain.UserResponse, []domain.ErrorMsg)
	DeleteUser(c *gin.Context) []domain.ErrorMsg
	UpdateUser(c *gin.Context) []domain.ErrorMsg
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(c *gin.Context) (domain.User, []domain.ErrorMsg) {
	return s.repository.CreateUser(c)
}

func (s *service) GetUser(c *gin.Context) (domain.UserResponse, []domain.ErrorMsg) {
	return s.repository.GetUser(c)
}

func (s *service) DeleteUser(c *gin.Context) []domain.ErrorMsg {
	return s.repository.DeleteUser(c)
}

func (s *service) UpdateUser(c *gin.Context) []domain.ErrorMsg {
	return s.repository.UpdateUser(c)
}
