package handlers

import (
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/gin-gonic/gin"
)

type User struct {
	userService user.Service
}

func NewUser(u user.Service) *User {
	return &User{
		userService: u,
	}
}

func Pong() func(c *gin.Context) {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (u *User) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := u.userService.CreateUser(c)
		c.JSON(http.StatusOK, user)
	}

}
