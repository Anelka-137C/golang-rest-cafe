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
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateUser(c *gin.Context) domain.User {
	return s.repository.CreateUser(c)
}
