package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
