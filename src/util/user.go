package util

import "github.com/gin-gonic/gin"

func BuildResponse(status int, object interface{}, message string, c *gin.Context) {
	if object != nil {
		c.JSON(status, gin.H{
			"msg":  message,
			"user": object,
		})
	} else {
		c.JSON(status, gin.H{
			"msg": message,
		})
	}
}
