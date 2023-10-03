package routes

import (
	controller"github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)


func OrderRoutes(incomingRoute *gin.Engine){
	incomingRoute.GET("/orders",controller.GetOrders())
	incomingRoute.GET("/orders/:order_id",controller.GetOneOrder())
	incomingRoute.POST("/orders",controller.CreateOrder())
	incomingRoute.PATCH("/orders/:order_id",controller.UpdateOrder())
}