package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

// InitMongo is a function that initializes a connection to the MongoDB database
func InitMongo() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URL must be set")
	}

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	client = conn
}

// CloseMongo is a function that closes the connection to the MongoDB database
func CloseMongo() {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

// UsersCollection is a function that returns the users collection
func UsersCollection() *mongo.Collection {
	return client.Database("infy").Collection("users")
}

// PostsCollection is a function that returns the posts collection
func PostsCollection() *mongo.Collection {
	return client.Database("infy").Collection("posts")
}

// CommentsCollection is a function that returns the comments collection
func CommentsCollection() *mongo.Collection {
	return client.Database("infy").Collection("comments")
}

// PollsCollection returns the polls collection from the database
func PollsCollection() *mongo.Collection {
	return client.Database("infy").Collection("polls")
}
