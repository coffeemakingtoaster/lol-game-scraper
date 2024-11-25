package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_NAME = "puuids"

var db *mongo.Database

func getMongoConnection() *mongo.Database {
	if db != nil {
		return db
	}
	godotenv.Load()

	uri := os.Getenv("MONGO_CONNECTION_STRING")
	conn, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	db = conn.Database("lol-scraper")
	ensureCollectionExists(db, COLLECTION_NAME)
	return db
}

func MarkPUUIDDone(puuid string) bool {
	return markPuuidDone(getMongoConnection(), puuid)
}

func markPuuidDone(db *mongo.Database, puuid string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection(COLLECTION_NAME)

	doc := bson.M{
		"puuid": puuid,
		"date":  time.Now(),
	}
	_, err := collection.InsertOne(ctx, doc)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func ensureCollectionExists(db *mongo.Database, collectionName string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the "puuids" collection exists
	exists, err := collectionExists(ctx, db, collectionName)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		fmt.Printf("Collection '%s' already exists.\n", collectionName)
	} else {
		fmt.Printf("Collection '%s' does not exist. Creating it...\n", collectionName)
		// To create a collection explicitly
		err = db.CreateCollection(ctx, collectionName)
		if err != nil {
			log.Fatalf("Failed to create collection '%s': %v\n", collectionName, err)
			panic(err)
		}
		collection := db.Collection(collectionName)

		// Ensure a unique index on the "puuid" field
		indexModel := mongo.IndexModel{
			Keys:    bson.M{"puuid": 1}, // Create a unique index on the "puuid" field
			Options: options.Index().SetUnique(true),
		}
		_, err = collection.Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Collection '%s' created successfully.\n", collectionName)
	}
}

func collectionExists(ctx context.Context, db *mongo.Database, collectionName string) (bool, error) {
	collections, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return false, err
	}

	for _, name := range collections {
		if name == collectionName {
			return true, nil
		}
	}
	return false, nil
}
