package httputil

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponseWithData(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"status": "failed",
		"data":   data,
	})
}

func ErrorResponseWithMessage(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
	})
}
