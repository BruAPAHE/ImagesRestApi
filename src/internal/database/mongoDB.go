package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ImagesConnection *mongo.Database

func init() {
	config := NewConfig()

	optionsDB := options.Client().SetMaxPoolSize(100)

	optionsDB.SetConnectTimeout(time.Minute)

	optionsDB.ApplyURI(config.DSN)

	client, err := mongo.NewClient(optionsDB)
	if err != nil {
		log.Printf("Cannot make connect: %v", err)
	}

	if err = client.Connect(context.TODO()); err != nil {
		log.Printf("Cannot init connect to MongoDB: %v", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	ImagesConnection = client.Database(config.Database)
}
