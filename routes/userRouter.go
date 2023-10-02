package routes

import (
	controller "github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/users", controller.GetUsers())
	incomingRoute.GET("/users/:user_id", controller.GetUser())
	incomingRoute.POST("/users/signup", controller.Signup())
	incomingRoute.POST("/users/login", controller.Login())
}
