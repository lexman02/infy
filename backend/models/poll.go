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

type Poll struct {
	ID        string    `bson:"_id" json:"id"`
	MovieID   string    `bson:"movie_id" json:"movie_id"`
	Question  string    `bson:"question" json:"question"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	EndsAt    time.Time `bson:"ends_at" json:"ends_at"`
	Options   []Option  `bson:"options" json:"options"`
}

type Option struct {
	ID    string `bson:"_id" json:"id"`
	Text  string `bson:"text" json:"text"`
	Votes int    `bson:"votes" json:"votes"`
}

func NewPoll(question, movieID string) *Poll {
	return &Poll{
		ID:        primitive.NewObjectID().Hex(),
		MovieID:   movieID,
		Question:  question,
		CreatedAt: time.Now(),
		EndsAt:    time.Now().Add(24 * time.Hour),
		Options:   nil,
	}
}

func (p *Poll) AddOption(text string) {
	p.Options = append(p.Options, Option{
		ID:    primitive.NewObjectID().Hex(),
		Text:  text,
		Votes: 0,
	})
}

func (p *Poll) Save(ctx context.Context) error {
	_, err := db.PollsCollection().InsertOne(ctx, p)
	if err != nil {
		return err
	}

	return nil
}

func IncrementPollOptionVote(pollID, optionID string, ctx context.Context) error {
	// Convert the pollID and optionID to ObjectIDs
	pollObjectID, err := primitive.ObjectIDFromHex(pollID)
	if err != nil {
		return err
	}

	// optionObjectID, err := primitive.ObjectIDFromHex(optionID)
	// if err != nil {
	// 	return err
	// }

	filter := bson.M{
		"_id": pollObjectID,
	}

	update := bson.M{
		"$inc": bson.M{"options.$[elem].votes": 1},
	}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem._id": optionID}},
	})

	poll, err := db.PollsCollection().UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		return err
	}

	fmt.Println("MatchedCount: ", poll.MatchedCount)
	fmt.Println("ModifiedCount: ", poll.ModifiedCount)
	fmt.Println("UpsertedCount: ", poll.UpsertedCount)
	fmt.Println("UpsertedID: ", poll.UpsertedID)

	return nil
}

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
