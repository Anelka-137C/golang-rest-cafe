package util

import (
	"github.com/Anelka-137C/cafe-app/internal/domain"
	"github.com/gin-gonic/gin"
)

func BuildResponse(status int, object interface{}, message string, c *gin.Context, entity string) {

	switch object.(type) {
	case domain.LoginResponse:
		c.JSON(status, gin.H{
			"token": object,
		})
		return
	default:
		if object != nil {
			c.JSON(status, gin.H{
				"msg":  message,
				entity: object,
			})
		} else {
			c.JSON(status, gin.H{
				"msg": message,
			})
		}
		return
	}

}

func BuildBadResponse(status int, err []domain.ErrorMsg, c *gin.Context) {
	c.JSON(status, gin.H{"error": err})
}