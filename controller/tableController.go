package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var tableCollection *mongo.Collection = database.OpenCollection(database.Client,"table")

func GetAllTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := tableCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when listing tables"})
			return
		}
		var allTables []bson.M
		err = result.All(ctx, &allTables)
		if err != nil {
			log.Fatal(err)
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func GetOneTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table model.Table
		tableId := c.Param("table_id")
		err = tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "table not found"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table model.Table
		err := c.BindJSON(&table)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validateErr := validate.Struct(table)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}
		table.Created_at = time.Now()
		table.Updated_at = time.Now()
		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()
		// table.Table_number =
		result, insertError := tableCollection.InsertOne(ctx, table)
		if insertError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "table creation failed"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
