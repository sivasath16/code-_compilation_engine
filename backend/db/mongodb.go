package backend

import (
	"context"
	"errors"
	"log"

	models "backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectMongoDB establishes a connection to MongoDB
func ConnectMongoDB() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
}

// SaveExecutionRequest saves a new execution request to the database
func SaveExecutionRequest(req models.CodeRequest) string {
	collection := client.Database("code_execution").Collection("tasks")
	result, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Fatalf("Failed to save request: %v", err)
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	log.Printf("Saved execution request with ID: %s", id)
	return id
}

// GetExecutionResult retrieves the execution result by ID
func GetExecutionResult(id string) (models.ExecutionResult, error) {
	collection := client.Database("code_execution").Collection("results")
	var result models.ExecutionResult

	log.Printf("Fetching result for Task ID: %s", id)

	// Query using the `id` field (string)
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result)
	if err != nil {
		log.Printf("Result not found for Task ID: %s", id)
		return result, errors.New("result not found")
	}

	log.Printf("Result found for Task ID: %s", id)
	return result, nil
}

// SaveExecutionResult saves the execution result in the database
func SaveExecutionResult(result models.ExecutionResult) {
	collection := client.Database("code_execution").Collection("results")

	// Ensure `id` field is populated
	if result.ID == "" {
		log.Printf("Execution result missing ID. Cannot save.")
		return
	}

	_, err := collection.InsertOne(context.TODO(), result)
	if err != nil {
		log.Printf("Failed to save execution result: %v", err)
	} else {
		log.Printf("Execution result saved with ID: %s", result.ID)
	}
}
