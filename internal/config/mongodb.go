package config

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient     *mongo.Client
	mongoClientOnce sync.Once
	dbName          = "NutriTrackAi"
)

// InitializeMongoClient initializes the MongoDB client if it doesn't exist
func InitializeMongoClient() {
	mongoClientOnce.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatal("Set your 'MONGODB_URI' environment variable. " +
				"See: www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}

		clientOpts := options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
		client, err := mongo.Connect(context.TODO(), clientOpts)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		if err := client.Ping(context.TODO(), nil); err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		mongoClient = client
		log.Println("Connected successfully to MongoDB")
	})
}

// GetMongoClient returns the initialized MongoDB client
func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		log.Fatal("MongoDB client is not initialized. Call InitializeMongoClient() first.")
	}
	return mongoClient
}

// GetCollection retrieves a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return GetMongoClient().Database(dbName).Collection(collectionName)
}

// InitializeCollections creates collections with optional schema validation
func InitializeCollections() {
	collections := []struct {
		name       string
		withSchema bool
		schema     bson.M
	}{
		{
			name:       "users",
			withSchema: true,
			schema: bson.M{
				"bsonType": "object",
				"required": []string{"name", "email", "age"},
				"properties": bson.M{
					"name":  bson.M{"bsonType": "string", "description": "must be a string"},
					"email": bson.M{"bsonType": "string", "description": "must be a string"},
					"age":   bson.M{"bsonType": "int", "minimum": 0, "description": "must be an integer greater than 0"},
				},
			},
		},
		{
			name:       "meals",
			withSchema: true,
			schema: bson.M{
				"bsonType": "object",
				"required": []string{"mealType", "name", "image", "date", "userId"},
				"properties": bson.M{
					"mealType": bson.M{"bsonType": "string", "description": "must be a string (e.g., breakfast, lunch)"},
					"name":     bson.M{"bsonType": "string", "description": "must be a string"},
					"image":    bson.M{"bsonType": "string", "description": "must be a string (URL or path to image)"},
					"date":     bson.M{"bsonType": "date", "description": "must be a valid date"},
					"userId":   bson.M{"bsonType": "string", "description": "must be a string (user's ID)"},
				},
			},
		},
		{
			name:       "labeled_images",
			withSchema: true,
			schema: bson.M{
				"bsonType": "object",
				"required": []string{"name", "image"},
				"properties": bson.M{
					"name":  bson.M{"bsonType": "string", "description": "must be a string"},
					"image": bson.M{"bsonType": "string", "description": "must be a string (URL or path to image)"},
				},
			},
		},
	}

	for _, col := range collections {
		if err := createCollection(col.name, col.withSchema, col.schema); err != nil {
			log.Printf("Error initializing collection '%s': %v\n", col.name, err)
		} else {
			log.Printf("Collection '%s' initialized successfully\n", col.name)
		}
	}
}

// Handles the creation of collections with optional schema validation
func createCollection(collectionName string, withSchema bool, schema bson.M) error {
	db := GetMongoClient().Database(dbName)
	collections, err := db.ListCollectionNames(context.TODO(), bson.M{"name": collectionName})
	if err != nil {
		return err
	}

	// If the collection already exists, nothing needs to be done
	if len(collections) > 0 {
		log.Printf("Collection '%s' already exists\n", collectionName)
		return nil
	}

	opts := options.CreateCollection()
	if withSchema && schema != nil {
		opts.SetValidator(bson.M{"$jsonSchema": schema})
	}

	return db.CreateCollection(context.TODO(), collectionName, opts)
}
