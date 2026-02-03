package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var Client *mongo.Client

func ConnectDatabase() {
	// Koneksi ke MongoDB
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		mongoURI = os.Getenv("MONGODB_URL")
	}

	if mongoURI == "" {
		// Use the Railway external URL provided by user as final fallback
		mongoURI = "mongodb://mongo:smVhjvhpEWPYtnCVKidYzgjIqMUBdcKm@gondola.proxy.rlwy.net:37059"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	DB = client.Database("rapatuy_penjaruy")
	Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
