package middleware

import (
	tokenJwt "go-commerce/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token not found",
			})
			c.Abort()
			return
		}

		res, err := tokenJwt.ValidateToken(token)
		if err != nil || res == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token invalid",
			})
			c.Abort()
			return
		}
		c.Set("user_id", res.ID)
		c.Next()
	}
}
