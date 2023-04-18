package router

import (
	"challenge-3-chapter-3/controllers"
	"challenge-3-chapter-3/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	roleRouter := router.Group("/role")
	{
		roleRouter.POST("/", controllers.CreateRole)
		roleRouter.GET("/", controllers.GetRole)
	}

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	productRouter := router.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/create", controllers.CreateProduct)
		productRouter.GET("all", controllers.GetAllProduct)
		productRouter.GET(":productID", middlewares.ProductAuthorization(), controllers.GetProductId)
		productRouter.PUT("/:productID", middlewares.RoleMiddleware(), middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productID", middlewares.RoleMiddleware(), middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return router
}
