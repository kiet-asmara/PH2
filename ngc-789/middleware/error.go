package middleware

import (
	"gin-ex/utils"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.ErrorMessage(c, &utils.ErrInternalServer)
			}
		}()

		c.Next() // make errormssg raw in the handler
		// only catch panics in middleware
	}
}
