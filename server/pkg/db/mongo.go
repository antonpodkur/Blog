package db

import (
	"context"
	"log"
	"time"

	"github.com/antonpodkur/Blog/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient(cfg *config.Config) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Mongo.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	return client
}

func OpenCollection(client *mongo.Client, cfg *config.Config, collectionName string) *mongo.Collection {
	collection := client.Database(cfg.Mongo.DbName).Collection(collectionName)
	return collection
}
