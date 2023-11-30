package repository

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type DB struct {
	client *mongo.Client
}

func NewDB() *DB {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &DB{client: client}
}

func (db *DB) GetCollection(databaseName string, collectionName string) *mongo.Collection {
	return db.client.Database(databaseName).Collection(collectionName)
}
