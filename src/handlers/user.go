package handlers

import (
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/gin-gonic/gin"
)

const (
	creationMessage = "User successfully created"
	getMessage      = "User successfully obtained"
	deleteMesage    = "User successfully delected"
	updateMessage   = "User successfully updated"
)

type User struct {
	userService user.Service
}

func NewUser(u user.Service) *User {
	return &User{
		userService: u,
	}
}

func (u *User) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := u.userService.CreateUser(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, user, creationMessage, c, "User")
	}
}

func (u *User) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := u.userService.GetUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		util.BuildResponse(http.StatusOK, user, getMessage, c, "User")
	}

}

func (u *User) GetUsersByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := u.userService.GetUsersByName(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		util.BuildResponse(http.StatusOK, users, getMessage, c, "Users")
	}

}

func (u *User) GetUserByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := u.userService.GetUserByEmail(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		util.BuildResponse(http.StatusOK, user, getMessage, c, "User")
	}

}

func (u *User) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := u.userService.DeleteUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		util.BuildResponse(http.StatusOK, nil, deleteMesage, c, "User")
	}

}

func (u *User) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := u.userService.UpdateUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		util.BuildResponse(http.StatusOK, nil, updateMessage, c, "User")
	}
}

func (u *User) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := u.userService.Login(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		util.BuildResponse(http.StatusOK, jwt, updateMessage, c, "User")
	}
}
