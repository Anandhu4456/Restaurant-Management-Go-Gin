package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/Anandhu4456/go-restaurant-management/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemPack struct {
	Table_id    *string
	Order_Items []model.OrderItem
}

var orderItemsCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItems")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result,err:=orderItemsCollection.Find(context.TODO(), bson.M{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when fetching order items"})
			return
		}
		var allOrderItem []bson.M
		err =result.All(ctx,&allOrderItem)
		if err!=nil{
			log.Fatal(err)
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,allOrderItem)
	}
}

func GetOneOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId:=c.Param("order_id")

		allOrders,err:=ItemsByOrderId(orderId)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when listing order items by id"})
			return
		}
		c.JSON(http.StatusOK,allOrders)
	}
}

func ItemsByOrderId(id string) (OrderItems []primitive.M, err error) {

}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
