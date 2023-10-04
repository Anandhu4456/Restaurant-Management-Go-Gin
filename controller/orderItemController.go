package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOneOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItemsByOrder()gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func ItemsByOrderId(id string) (OrderItems []primitive.M, err error) {

}

func CreateOrderItem()gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateOrderItem()gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}
