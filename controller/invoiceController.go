package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Order_details    interface{}
	Payment_due_date time.Time
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result,err:=invoiceCollection.Find(context.TODO(), bson.M{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when fetching invoices"})
			return
		}
		var allInvoices []bson.M
		err =result.All(ctx,&allInvoices)
		if err!=nil{
			log.Fatal(err)
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
	}
}

func GetOneInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
