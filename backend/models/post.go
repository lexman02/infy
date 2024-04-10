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
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	User      *User              `json:"user" bson:"user"`
	Reactions []UserReactions    `json:"-"`
	Movie     *Movie             `json:"movie"`
	Content   string             `json:"content"`
}

type UserReactions struct {
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Liked    bool               `json:"liked"`
	Disliked bool               `json:"disliked"`
}

type Movie struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	PosterPath string `json:"poster_path"`
	Tagline    string `json:"tagline"`
}

// NewPost creates a new post instance
func NewPost(user *User, movie *Movie, content string) *Post {
	return &Post{ID: primitive.NewObjectID(), User: user, Reactions: nil, Movie: movie, Content: content}
}

// FindAllPosts finds all the posts
func FindAllPosts(ctx context.Context, limit int64) ([]*Post, error) {
	opts := options.Find().SetSort(bson.D{bson.E{Key: "createdAt", Value: -1}}).SetLimit(limit) // Corrected line for sorting
	cursor, err := db.PostsCollection().Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*Post
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

func UpdateReaction(postID string, userID primitive.ObjectID, like, dislike bool, ctx context.Context) error {
	// Encode the post ID to an ObjectID type
	postObjectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	// Create a UserReactions object
	reaction := UserReactions{UserID: userID, Liked: like, Disliked: dislike}

	// Remove the existing reaction from the user
	_, err = db.PostsCollection().UpdateOne(ctx, bson.M{"_id": postObjectID}, bson.M{"$pull": bson.M{"reactions": bson.M{"user_id": userID}}})
	if err != nil {
		return err
	}

	// Add the new reaction from the user
	_, err = db.PostsCollection().UpdateOne(ctx, bson.M{"_id": postObjectID}, bson.M{"$push": bson.M{"reactions": reaction}})
	if err != nil {
		return err
	}

	return nil
}

func FindPostsByUserID(userID string, ctx context.Context, limit int64) ([]*Post, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Define sorting and limiting options
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetLimit(limit)

	filter := bson.M{"user._id": userObjectID}
	cursor, err := db.PostsCollection().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []*Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func FindPostsByMovieID(movieID string, ctx context.Context, limit int64) ([]*Post, error) {
	var posts []*Post

	// Define sorting and limiting options
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetLimit(limit)

	filter := bson.M{"movie_id": movieID}
	cursor, err := db.PostsCollection().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
