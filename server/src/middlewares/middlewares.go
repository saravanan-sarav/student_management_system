package middlewares

import "github.com/gin-gonic/gin"


func WithAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Next()
	}
}