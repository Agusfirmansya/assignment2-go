package routers

import (
	"test_assignment2/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/orders", controllers.GetAllOrder)
	router.GET("/orders/:orderID", controllers.GetOrder)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)


	return router
}