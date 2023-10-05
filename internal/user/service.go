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
	CreateUser(c *gin.Context) domain.User
	GetUser(c *gin.Context) domain.User
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(c *gin.Context) domain.User {
	return s.repository.CreateUser(c)
}

func (s *service) GetUser(c *gin.Context) domain.User {
	return s.repository.GetUser(c)
}

func (s *service) DeleteUser(c *gin.Context) {
	s.repository.DeleteUser(c)
}

func (s *service) UpdateUser(c *gin.Context) {
	s.repository.UpdateUser(c)
}
