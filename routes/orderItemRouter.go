package routes

import (
	controller "github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/orderItems", controller.GetOrderItems())
	incomingRoute.GET("/orderItems/:orderItem_id", controller.GetOneOrderItem())
	incomingRoute.GET("/orderItems-order", controller.GetOrderItemsByOrder())
	incomingRoute.POST("/orderItems", controller.CreateOrderItem())
	incomingRoute.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
