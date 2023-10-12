package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Anandhu4456/go-restaurant-management/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err2 := strconv.Atoi(c.Query("startIndex"))
		if err2 != nil {
			// error
		}

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		projectStage := bson.D{
			{
				Key:"$project",Value: bson.D{
					{Key:"_id",Value: 0},
					{Key:"total_count",Value:  1},
					{Key:"user_items",Value:  bson.D{
						{Key:"$slice",Value: bson.A{"$data", startIndex, recordPerPage}},
					}},
				},
			},
		}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, projectStage})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured when listing user items"})
			return
		}
		var allUsers[]bson.M
		if err:=result.All(ctx,&allUsers);err!=nil{
			log.Fatal(err)
		}
		defer cancel()
		c.JSON(http.StatusOK,allUsers[0])
	}

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) string {

}

func VerifyPassword(userPassword, providedPassword string) (bool, string) {

}
