package handlers

import (
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/domain"
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

func (u *User) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := u.userService.GetUser(c)
		c.JSON(http.StatusOK, user)
	}

}

func (u *User) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u.userService.DeleteUser(c)

		c.JSON(http.StatusOK, domain.Message{
			Msg: "Documen deleted",
		})
	}

}

func (u *User) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		u.userService.UpdateUser(c)

		c.JSON(http.StatusOK, domain.Message{
			Msg: "Documen updated",
		})
	}

}
