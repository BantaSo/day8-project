package route

import (
	"day8-project/controller"

	"github.com/gin-gonic/gin"
)

func StartRoute() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controller.TakeItem)
	router.DELETE("/order/orderId", controller.TakeDown)
	router.GET("/orders", controller.TakeItem)
	router.PUT("/order/orderID", controller.UpdateItem)
	return router
}
