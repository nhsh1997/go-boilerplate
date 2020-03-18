package dbs

import (
	"context"
	"fmt"
	configs "go-boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
type MongoDBClient struct {
	Client *mongo.Client

}

func NewMongoDBClient (configuration configs.Configuration) *MongoDBClient{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoDBClient{
		Client: client,
	}
}