package middlewares

import (
	"challenge-3-chapter-3/database"
	"challenge-3-chapter-3/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		productID, _ := strconv.Atoi(ctx.Param("productID"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		roleID := uint(userData["role"].(float64))

		if roleID == 1 {
			ctx.Next()
		} else {
			var product models.Product
			err := db.Select("user_id").First(&product, uint(productID)).Error
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   "Data not found",
					"message": "data doesn't exist",
				})
				return
			}
			if product.UserID != userID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			} else {
				ctx.Next()
			}
		}
	}
}
