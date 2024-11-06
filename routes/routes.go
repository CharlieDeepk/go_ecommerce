package routes

import (
	"github.com/CharlieDeepk/go_ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.LogIn())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("users/productview", controllers.SearchProduct())
	incomingRoutes.GET("users/search", controllers.SearchProductByQuery())
}
