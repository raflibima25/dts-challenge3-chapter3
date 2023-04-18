package controllers

import (
	"challenge-3-chapter-3/database"
	"challenge-3-chapter-3/helpers"
	"challenge-3-chapter-3/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	var product models.Product

	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetProductId(c *gin.Context) {
	db := database.GetDB()
	var product models.Product

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	err = db.Debug().First(&product, "id = ?", productID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	var product []models.Product

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	roleID := uint(userData["role"].(float64))

	if roleID == 1 {
		err := db.Debug().Find(&product).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if roleID == 2 {
		err := db.Debug().Where("user_id", userID).Find(&product).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data_product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	var product, findProduct models.Product

	contentType := helpers.GetContentType(c)

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product = models.Product{
		Title:       product.Title,
		Description: product.Description,
	}

	err = db.Where("id = ?", productID).First(&findProduct).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	product.ID = uint(productID)
	product.UserID = findProduct.UserID

	err = db.Model(&product).Where("id = ?", productID).Updates(product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	var product models.Product

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		log.Println("error di product ID")
		return
	}

	err = db.Debug().Where("id = ?", productID).First(&product).Delete(&product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product %s success deleted", product.Title),
	})
}
