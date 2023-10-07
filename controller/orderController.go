package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/Anandhu4456/go-restaurant-management/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result,err:=orderCollection.Find(context.TODO(),bson.M{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		var allOrders []bson.M
		defer cancel()
		err = result.All(ctx,&allOrders)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when listing order items"})
			return
		}
		c.JSON(http.StatusOK,allOrders)
	}
}

func GetOneOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel =context.WithTimeout(context.Background(),100*time.Second)
		orderId:=c.Param("order_id")
		var order model.Order
		err:=orderCollection.FindOne(ctx,bson.M{"order_id":orderId}).Decode(&order)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"order not found"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
