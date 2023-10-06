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

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when listing the menu items"})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetOneMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu model.Menu
		menuId := c.Param("menu_id")
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured when fetching the menu"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu model.Menu
		err:=c.BindJSON(menu)
		if err!=nil{
			c.JSON(http.StatusBadRequest,bson.M{"error":err.Error()})
			return
		}
		
		validateError:=validate.Struct(menu)
		if validateError!=nil{
			c.JSON(http.StatusBadRequest,bson.M{"error":err.Error()})
			return
		}

		menu.Created_at =time.Now()
		menu.Updated_at = time.Now()
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()
		
		result,insertErr:=menuCollection.InsertOne(ctx,menu)
		if insertErr!=nil{
			c.JSON(http.StatusInternalServerError,bson.M{"error":"error while creating menu"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK,result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
