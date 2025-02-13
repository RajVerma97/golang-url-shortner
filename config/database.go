package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client
var Database *mongo.Database
var Collection *mongo.Collection

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found ")

	}

	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {
		log.Fatal("MONGODB_URI IS NOT DEFINED ")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal("MONGODB CONNECTION ERROR: ", err)
	}

	Client = client
	Database = Client.Database("url-shortner")
	Collection = Database.Collection("urls")
	log.Println("Successfully connnected to the Mongodb")
}
