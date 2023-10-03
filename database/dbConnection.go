package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb....")
	return client
}

var Client *mongo.Client = DbInstance()

func OpenCollection(client *mongo.Client,collectionName string)*mongo.Collection{
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
