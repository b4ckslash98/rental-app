package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin only"})
			return
		}
		c.Next()
	}
}

func CustomerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != "customer" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Customer only"})
			return
		}
		c.Next()
	}
}
