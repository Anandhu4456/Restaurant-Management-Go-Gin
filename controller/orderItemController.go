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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderItemPack struct {
	Table_id    *string
	Order_Items []model.OrderItem
}

var orderItemsCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItems")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemsCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when fetching order items"})
			return
		}
		var allOrderItem []bson.M
		err = result.All(ctx, &allOrderItem)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, allOrderItem)
	}
}

func GetOneOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var orderItemId = c.Param("order_item_id")
		var orderItem model.OrderItem
		err := orderItemsCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId}).Decode(&orderItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "couldnt get order item id"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, orderItem)
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")

		allOrders, err := ItemsByOrderId(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when listing order items by id"})
			return
		}
		c.JSON(http.StatusOK, allOrders)
	}
}

func ItemsByOrderId(id string) (OrderItems []primitive.M, err error) {
	var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)

	matchStage:=bson.D{{"$match",bson.D{{"order_id",id}}}}
	lookupStage:= bson.D{{"$lookup",bson.D{{"from","food"},{"localField","food_id",{"foreignFiled","food_id"},{"as","food"}}}}}
	unwindStage:= bson.D{{"$unwind",bson.D{{"path","$food"},{"preserveNullAndEmptyArrays",true}}}}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// var orderItem model.OrderItem
		var order model.Order
		var orderItemPack OrderItemPack

		err := c.BindJSON(&orderItemPack)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		order.Order_date = time.Now().AddDate(0, 0, 1)
		orderItemsToBeInserted := []interface{}{}
		order.Table_id = orderItemPack.Table_id
		order.Order_id = OrderItemsOrderCreater(order)

		for _, orderItem := range orderItemPack.Order_Items {
			orderItem.Order_id = order.Order_id

			validateErr := validate.Struct(orderItem)
			if validateErr != nil {
				c.JSON(http.StatusBadRequest,gin.H{"error":validateErr.Error()})
				return
			}
			orderItem.ID = primitive.NewObjectID()
			orderItem.Created_at = time.Now()
			orderItem.Updated_at = time.Now()
			orderItem.Order_item_id = orderItem.ID.Hex()

			var num = toFixed(*orderItem.Unit_price,2)
			orderItem.Unit_price = &num
			orderItemsToBeInserted = append(orderItemsToBeInserted, orderItem)
		}
		insertedOrderItem,err:=orderItemsCollection.InsertMany(ctx,orderItemsToBeInserted)
		if err!=nil{
			log.Fatal(err)
		}
		defer cancel()
		c.JSON(http.StatusOK,insertedOrderItem)
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
		var orderItem model.OrderItem
		orderItemId:=c.Param("order_item_id")

		filter:=bson.M{"order_item_id":orderItemId}

		err:=c.BindJSON(&orderItem)
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}
		var updateObj primitive.D

		if orderItem.Unit_price!=nil{
			updateObj = append(updateObj,bson.E{Key: "unit_price", Value: orderItem.Unit_price})
		}
		if orderItem.Quantity !=nil{
			updateObj = append(updateObj, bson.E{Key: "quantity",Value: orderItem.Quantity})
		}
		if orderItem.Food_id !=nil{
			updateObj = append(updateObj, bson.E{Key: "food_id",Value: orderItem.Food_id})
		}
		orderItem.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at",Value: orderItem.Updated_at})

		upsert:=true

		opt:=options.UpdateOptions{
			Upsert: &upsert,
		}
		result,updateError:=orderItemsCollection.UpdateOne(
			ctx,
			filter,
			primitive.D{
				{Key:"$set",Value: updateObj},
			},
			&opt,
		)
		if updateError!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when updaing order items"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
	}
}
