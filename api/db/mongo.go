package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(uri, name string) (*mongo.Database, error) {
	clientOpts := options.Client().ApplyURI(uri)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Println("Error on connect to MongoDB:", err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Mongo connection error:", err)
		return nil, err
	}
	// defer client.Disconnect(ctx)
	return client.Database(name), nil
}
