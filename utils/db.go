package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDB(ctx context.Context) *mongo.Client {
mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return mongoclient
}