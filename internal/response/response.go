package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, status int, message string, data any) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func SuccessMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": true,
		"message": message,
	})
}

func Error(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func ErrorMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"error":   message,
	})
}
