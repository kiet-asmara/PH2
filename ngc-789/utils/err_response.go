package utils

import "github.com/gin-gonic/gin"

func ErrorMessage(c *gin.Context, apiError *APIError) *gin.Context {
	c.Abort()
	c.JSON(apiError.Code, gin.H{"error": APIError{
		Code:    apiError.Code,
		Message: apiError.Message,
	}})
	return c
}
