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
		user, err := u.userService.CreateUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "User created",
			"user": user,
		})
	}

}

func (u *User) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := u.userService.GetUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "User successfully obtained ",
			"user": user,
		})
	}

}

func (u *User) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := u.userService.DeleteUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, domain.Message{
			Msg: "User deleted",
		})
	}

}

func (u *User) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := u.userService.UpdateUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, domain.Message{
			Msg: "User updated",
		})
	}

}
