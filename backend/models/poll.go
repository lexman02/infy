package models

import (
	"context"
	"fmt"
	"infy/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Poll represents a poll related to a movie, containing multiple options for responses.
type Poll struct {
	ID        string    `bson:"_id" json:"id"`
	MovieID   string    `bson:"movie_id" json:"movie_id"`
	Question  string    `bson:"question" json:"question"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	EndsAt    time.Time `bson:"ends_at" json:"ends_at"`
	Options   []Option  `bson:"options" json:"options"`
}

// Option defines a single selectable option within a poll.
type Option struct {
	ID    string `bson:"_id" json:"id"`
	Text  string `bson:"text" json:"text"`
	Votes int    `bson:"votes" json:"votes"`
}

// NewPoll creates a new Poll with the specified question and associated movie ID.
func NewPoll(question, movieID string) *Poll {
	return &Poll{
		ID:        primitive.NewObjectID().Hex(),
		MovieID:   movieID,
		Question:  question,
		CreatedAt: time.Now(),
		EndsAt:    time.Now().Add(24 * time.Hour), // Poll ends in 24 hours from creation
		Options:   nil,
	}
}

// AddOption adds a new option to the Poll.
func (p *Poll) AddOption(text string) {
	p.Options = append(p.Options, Option{
		ID:    primitive.NewObjectID().Hex(),
		Text:  text,
		Votes: 0, // Initialize votes to zero for the new option
	})
}

// Save inserts the Poll into the database.
func (p *Poll) Save(ctx context.Context) error {
	_, err := db.PollsCollection().InsertOne(ctx, p)
	return err
}

// IncrementPollOptionVote increases the vote count for a specific option in a poll.
func IncrementPollOptionVote(pollID, optionID string, ctx context.Context) error {
	pollObjectID, err := primitive.ObjectIDFromHex(pollID)
	if err != nil {
		return err // Handle invalid ObjectID format
	}

	filter := bson.M{"_id": pollObjectID}
	update := bson.M{"$inc": bson.M{"options.$[elem].votes": 1}}
	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem._id": optionID}},
	})

	result, err := db.PollsCollection().UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		return err
	}

	fmt.Println("MatchedCount: ", result.MatchedCount)
	fmt.Println("ModifiedCount: ", result.ModifiedCount)
	fmt.Println("UpsertedCount: ", result.UpsertedCount)
	fmt.Println("UpsertedID: ", result.UpsertedID)

	return nil
}

// FindPollsByMovieID retrieves all polls associated with a specific movie ID.
func FindPollsByMovieID(movieID string, ctx context.Context) ([]*Poll, error) {
	var polls []*Poll
	filter := bson.M{"movie_id": movieID}
	cursor, err := db.PollsCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &polls); err != nil {
		return nil, err
	}

	return polls, nil
}
