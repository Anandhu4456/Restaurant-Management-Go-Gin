package routes

import (
	controller "github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/foods", controller.GetAllFoods())
	incomingRoute.GET("/foods/:food_id", controller.GetOneFoodItem())
	incomingRoute.POST("/foods", controller.CreateFood())
	incomingRoute.PATCH("/foods/:food_id", controller.UpdateFood())
}
