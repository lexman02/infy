package models

import (
	"context"
	"infy/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	User     *User              `json:"user" bson:"user"`
	Likes    int                `json:"likes"`
	Dislikes int                `json:"dislikes"`
	Content  string             `json:"content"`
}

// NewPost creates a new post instance
func NewPost(user *User, content string) *Post {
	return &Post{ID: primitive.NewObjectID(), User: user, Likes: 0, Dislikes: 0, Content: content}
}

// FindAllPosts finds all the posts
func FindAllPosts(ctx context.Context) ([]*Post, error) {
	var posts []*Post

	// Find all posts
	cursor, err := db.PostsCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Decode all the posts
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

// FindPostByID finds a post by ID
func FindPostByID(id string, ctx context.Context) (*Post, error) {
	var post Post

	// Encode the ID to an ObjectID type
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the post by ID
	err = db.PostsCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&post)

	return &post, err
}

// UpdateUserPost updates a post by ID and user ID
func UpdateUserPost(id, content string, userId primitive.ObjectID, ctx context.Context) error {
	// Encode the ID to an ObjectID type
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Create the filter and update
	filter := bson.D{{Key: "_id", Value: objectID}, {Key: "user._id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "content", Value: content}}}}

	// Set the return document to after
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Find the post by ID and user ID and update it if they match
	var post Post
	err = db.PostsCollection().FindOneAndUpdate(ctx, filter, update, opts).Decode(&post)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserPost deletes a post by ID and user ID
func DeleteUserPost(id string, userId primitive.ObjectID, ctx context.Context) error {
	// Encode the ID to an ObjectID type
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Create the filter
	filter := bson.D{{Key: "_id", Value: objectID}, {Key: "user._id", Value: userId}}

	// Delete the post by ID and user ID if they match
	results, err := db.PostsCollection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if nothing was deleted and return an error
	if results.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil

}

// Save saves a post to the database
func (p *Post) Save(ctx context.Context) error {
	// Insert the post into the database
	_, err := db.PostsCollection().InsertOne(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

// AddLike adds a like to the post
func (p *Post) AddLike() {
	p.Likes++
}

// AddDislike adds a dislike to the post
func (p *Post) AddDislike() {
	p.Dislikes++
}
