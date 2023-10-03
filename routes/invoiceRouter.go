package routes

import (
	controller "github.com/Anandhu4456/go-restaurant-management/controller"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoute *gin.Engine){
	incomingRoute.GET("/invoices",controller.GetInvoices())
	incomingRoute.GET("/invoices/:invoice_id",controller.GetOneInvoice())
	incomingRoute.POST("/invoices",controller.CreateInvoice())
	incomingRoute.PATCH("invoices/:invoice_id",controller.UpdateInvoice())
}
