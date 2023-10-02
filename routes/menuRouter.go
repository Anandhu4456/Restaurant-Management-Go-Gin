package routes

import (
	controller "github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoute *gin.Engine) {
	incomingRoute.GET("/menus", controller.GetMenus())
	incomingRoute.GET("/menus/:menu_id", controller.GetOneMenu())
	incomingRoute.POST("/menus/createmenu",controller.CreateMenu())
	incomingRoute.PATCH("menus/:menu_id",controller.UpdateMenu())

}
