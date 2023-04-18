package controllers

import (
	"challenge-3-chapter-3/database"
	"challenge-3-chapter-3/helpers"
	"challenge-3-chapter-3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func RegisterUser(c *gin.Context) {
	db := database.GetDB()
	var user models.User
	var role models.Role

	contentType := helpers.GetContentType(c)

	if contentType == appJson {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	err := db.Debug().First(&role, user.RoleID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Role ID not found",
		})
		return
	}

	err = db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"full_name": user.FullName,
			"email":     user.Email,
			"role":      role.RoleName,
		},
		"message": "User success created",
	})
}

func LoginUser(c *gin.Context) {
	db := database.GetDB()
	var user models.User

	contentType := helpers.GetContentType(c)

	if contentType == appJson {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.RoleID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
