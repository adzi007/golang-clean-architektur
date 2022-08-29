package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbConnect struct{}

type mongoDbInterface interface {
	ConnectDB() *mongo.Client
}

var MongoDBConfig mongoDbInterface = MongoDbConnect{}

func (dbCon MongoDbConnect) ConnectDB() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:kopiluwak01@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var DBConnection *mongo.Client = MongoDBConfig.ConnectDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("quickstart").Collection(collectionName)
	return collection
}
