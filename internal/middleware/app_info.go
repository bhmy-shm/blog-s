package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Blog_name", "goweb-service")
		c.Set("Blog_version", "1.0.0")
		c.Next()
	}
}
