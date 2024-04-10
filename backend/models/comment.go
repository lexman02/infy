package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"infy/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	ID         primitive.ObjectID   `json:"id" bson:"_id"`
	PostID     primitive.ObjectID   `json:"post_id" bson:"post_id"`
	User       *User                `json:"user" bson:"user"`
	Likes      int                  `json:"likes"`
	Dislikes   int                  `json:"dislikes"`
	LikedBy    []primitive.ObjectID `bson:"liked_by" json:"liked_by,omitempty"`
	DislikedBy []primitive.ObjectID `bson:"disliked_by" json:"disliked_by,omitempty"`
	Content    string               `json:"content"`
}

// NewComment creates a new comment instance
func NewComment(postID primitive.ObjectID, user *User, content string) *Comment {
	return &Comment{ID: primitive.NewObjectID(), PostID: postID, User: user, Likes: 0, Dislikes: 0, Content: content}
}

//FindCommentsByUser finds all comments created by a user
//TODO if we want to implement

// FindCommentByID finds comments by comment ID
func FindCommentsByID(id string, ctx context.Context) (*Comment, error) {
	var comment Comment

	// Encode the ID to an ObjectID type
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find the post by ID
	err = db.CommentsCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&comment)

	return &comment, err
}

/*

Next function might be redundant so idk if we want to just keep one saying findcommentbyid

*/

// FindCommentsByPostID finds comments by post ID
func FindCommentsByPostID(postID string, ctx context.Context, limit int64) ([]*Comment, error) {
	postObjectID, _ := primitive.ObjectIDFromHex(postID)                                        // Assuming postID is valid and converting it to ObjectID
	opts := options.Find().SetSort(bson.D{bson.E{Key: "createdAt", Value: -1}}).SetLimit(limit) // Sorting by createdAt in descending order
	cursor, err := db.CommentsCollection().Find(ctx, bson.M{"postId": postObjectID}, opts)
	if err != nil {
		return nil, err
	}
	var comments []*Comment
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// DeleteUserComment deletes a comment by ID and user ID
func DeleteUserComment(id string, userId primitive.ObjectID, ctx context.Context) error {
	// Encode the ID to an ObjectID type
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Create the filter
	filter := bson.D{{Key: "_id", Value: objectID}, {Key: "user._id", Value: userId}}

	// Delete the post by ID and user ID if they match
	results, err := db.CommentsCollection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if nothing was deleted and return an error
	if results.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil

}

// UpdateUserComment updates a comment by ID and user ID
func UpdateUserComment(id, content string, userId primitive.ObjectID, ctx context.Context) error {
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

	// Find the comment by ID and user ID and update it
	var comment Comment
	err = db.CommentsCollection().FindOneAndUpdate(ctx, filter, update, opts).Decode(&comment)
	if err != nil {
		return err
	}

	return nil
}

func (c *Comment) Save(ctx context.Context) error {
	// Insert the post into the database
	_, err := db.CommentsCollection().InsertOne(ctx, c)
	if err != nil {
		return err
	}

	return nil
}

// AddLike adds a like to the comment
func (c *Comment) AddLike() {
	c.Likes++
}

// AddDislike adds a dislike to the comment
func (c *Comment) AddDislike() {
	c.Dislikes++
}

// ToggleLikeOnComment updates the like or dislike state of a comment for a given user
func ToggleLikeOnComment(commentID string, userID string, like bool) error {

	cID, _ := primitive.ObjectIDFromHex(commentID)
	uID, _ := primitive.ObjectIDFromHex(userID)

	update := bson.M{}
	if like {
		update = bson.M{
			"$addToSet": bson.M{"liked_by": uID},
			"$pull":     bson.M{"disliked_by": uID},
		}
	} else {
		update = bson.M{
			"$addToSet": bson.M{"disliked_by": uID},
			"$pull":     bson.M{"liked_by": uID},
		}
	}

	_, err := db.CommentsCollection().UpdateOne(context.Background(), bson.M{"_id": cID}, update)
	return err
}
