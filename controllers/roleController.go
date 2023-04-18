package controllers

import (
	"challenge-3-chapter-3/database"
	"challenge-3-chapter-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	db := database.GetDB()
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Debug().Create(&role).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid create data",
		})
		return
	}

	c.JSON(http.StatusCreated, role)
}

func GetRole(c *gin.Context) {
	db := database.GetDB()
	var roleDatas []models.Role

	err := db.Debug().Find(&roleDatas).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": roleDatas,
	})

}
