package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userData := c.MustGet("userData").(jwt.MapClaims)
		roleID := uint(userData["role"].(float64))

		if roleID != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not admin",
			})
			return
		}

		c.Next()
	}
}
