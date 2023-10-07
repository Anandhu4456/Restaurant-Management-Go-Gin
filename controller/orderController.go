package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/Anandhu4456/go-restaurant-management/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var tableCollection *mongo.Collection = database.OpenCollection(database.Client,"table")
var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
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
		// var ctx,cancel =context.WithTimeout(context.Background(),100*time.Second)
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
		// var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

		var table model.Table
		var order model.Order
		err:=c.BindJSON(&order)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		err =validate.Struct(order)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		if order.Table_id!=nil{
			err:=tableCollection.FindOne(ctx,bson.M{"table_id":order.Table_id}).Decode(&table)
			if err!=nil{
				c.JSON(http.StatusInternalServerError,gin.H{"error":"table not found"})
				return
			}
		}
		order.Created_at = time.Now()
		order.Updated_at = time.Now()
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result,createError:=orderCollection.InsertOne(ctx,order)
		if createError!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"order creation faild"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

		orderId:=c.Param("order_id")
		var order model.Order
		var table model.Table
		err:=c.BindJSON(&order)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		var updateObj primitive.D

		if order.Table_id!=nil{
		err:=tableCollection.FindOne(ctx,bson.M{"table_id":order.Table_id}).Decode(&table)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"could't find table"})
			return
		}
		updateObj = append(updateObj, bson.E{Key:"table_id",Value:order.Table_id})
		}
		order.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key:"updated_at",Value:order.Updated_at})

		upsert:=true
		filter:=bson.M{"order_id":orderId}
		
		opt:=options.UpdateOptions{
			Upsert: &upsert,
		}
		result,updateError:=orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set",Value: updateObj},
			},
			&opt,
		)
		if updateError!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"couldn't update order"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
		
	}
}

func OrderItemsOrderCreater(order model.Order)string{
	order.Created_at =time.Now()
	order.Updated_at = time.Now()
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx,order)
	return order.Order_id
}