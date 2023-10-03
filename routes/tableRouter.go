package routes

import (
	controller"github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoute *gin.Engine){
	incomingRoute.GET("/tables",controller.GetAllTables())
	incomingRoute.GET("/tables/:table_id",controller.GetOneTable())
	incomingRoute.POST("/tables",controller.CreateTable())
	incomingRoute.PATCH("/tables/:table_id",controller.UpdateTable())
}
