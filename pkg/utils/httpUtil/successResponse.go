package httputil

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponseWithData(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"status": "success",
		"data":   data,
	})
}

func SuccessResponseWithMessage(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
	})
}
