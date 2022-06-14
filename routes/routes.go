package routes

import(
	"github.com/vikas-gautam/ecommerce-cart/controllers" 
    "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductviewerAdmin()) //in future we could add authorization 
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())

}