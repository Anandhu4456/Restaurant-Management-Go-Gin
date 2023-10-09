package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// var tableCollection *mongo.Collection = database.OpenCollection(database.Client,"table")

func GetAllTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result,err:=tableCollection.Find(context.TODO(), bson.M{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when listing tables"})
			return
		}
		var allTables []bson.M
		err = result.All(ctx,&allTables)
		if err!=nil{
			log.Fatal(err)
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
	}
}

func GetOneTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
