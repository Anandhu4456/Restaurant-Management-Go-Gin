package controller

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/Anandhu4456/go-restaurant-management/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var err error
var validate = validator.New()

func GetAllFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		filterStage := bson.D{{Key: "$group",Value: bson.D{
			{Key: "_id",Value:"null"},
			{Key: "total_count",Value: bson.D{{Key: "$sum",Value: "1"}}},
			{Key: "data",Value: bson.D{{Key: "$push",Value: "$$ROOT"}}},
		}}}
		projectStage := bson.D{
			{
				Key: "$project", Value: bson.D{
					{Key: "_id", Value: 0},
					{Key: "total_count", Value: 1},
					{Key: "food_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}},
				},
			},
		}
		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, filterStage, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when listing food items"})
			return
		}
		var allFoods []bson.M
		err = result.All(ctx, &allFoods)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allFoods[0])
	}
}

func GetOneFoodItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food model.Food
		foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when fetching the food item"})
			return
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu model.Menu
		var food model.Food

		if err = c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationError := validate.Struct(food)
		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			msg := fmt.Sprintf("menu not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		food.Created_at = time.Now()
		food.Updated_at = time.Now()
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, insertError := foodCollection.InsertOne(ctx, food)
		if insertError != nil {
			msg := fmt.Sprintf("food was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {
	return int(num+math.Copysign(0.5,num))
}

func toFixed(num float64, precision int) float64 {
	output:=math.Pow(10,float64(precision))
	return float64(round(num*output))/output
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var food model.Food
		var menu model.Menu

		foodID := c.Param("food_id")
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D
		if food.Name != nil {
			updateObj = append(updateObj, bson.E{Key: "name", Value: food.Name})
		}
		if food.Price != nil {
			updateObj = append(updateObj, bson.E{Key: "price", Value: food.Price})
		}
		if food.Food_image != nil {
			updateObj = append(updateObj, bson.E{Key: "food_image", Value: food.Food_image})
		}
		if food.Menu_id != nil {
			err := menuCollection.FindOne(ctx,bson.M{"menu_id": food.Menu_id}).Decode(&menu)
			defer cancel()
			if err!=nil{
				c.JSON(http.StatusInternalServerError,gin.H{"error":"menu not found"})
			}
			updateObj = append(updateObj, bson.E{Key: "menu", Value:food.Price})
		}

		food.Updated_at = time.Now()
		updateObj = append(updateObj, bson.E{Key: "updated_at",Value: food.Updated_at})

		upsert := true
		filter := bson.M{"food_id": foodID}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, updateErr := foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "food item not updated"})
		}
		c.JSON(http.StatusOK, result)
	}
}
