package models

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var (
	defaultUser   = NewUser("", "test@example.com", "", NewProfile("", "", time.Time{}, NewPreferences()))
	expectedPosts = []*Post{
		{
			ID:        primitive.NewObjectID(),
			User:      defaultUser,
			Reactions: nil,
			Movie: &Movie{
				ID:         1,
				Title:      "Test Movie",
				PosterPath: "",
				Tagline:    "Test Tagline",
			},
			Content: "Test Content",
		},
		{
			ID:        primitive.NewObjectID(),
			User:      defaultUser,
			Reactions: nil,
			Movie: &Movie{
				ID:         1,
				Title:      "Test Movie",
				PosterPath: "",
				Tagline:    "Test Tagline",
			},
			Content: "Test Content",
		},
		{
			ID:        primitive.NewObjectID(),
			User:      defaultUser,
			Reactions: nil,
			Movie: &Movie{
				ID:         1,
				Title:      "Test Movie",
				PosterPath: "",
				Tagline:    "Test Tagline",
			},
			Content: "Test Content",
		},
	}
)

func TestNewPost(t *testing.T) {
	user := &User{Email: "test@example.com"}
	movie := &Movie{ID: 1, Title: "Test Movie", PosterPath: "/path/to/poster", Tagline: "Test Tagline"}
	content := "Test Content"

	post := NewPost(user, movie, content)

	assert.NotNil(t, post)
	assert.NotEmpty(t, post.ID)
	assert.Equal(t, user, post.User)
	assert.Nil(t, post.Reactions)
	assert.Equal(t, movie, post.Movie)
	assert.Equal(t, content, post.Content)
}

func TestFindAllPosts(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("find all posts", func(mt *mtest.T) {
		// Create a cursor response with the expected posts
		ns := mt.DB.Name() + "." + mt.Coll.Name()
		first := mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedPosts[0].ID},
			{Key: "user", Value: defaultUser},
			{Key: "movie", Value: bson.D{{Key: "id", Value: expectedPosts[0].Movie.ID}, {Key: "title", Value: expectedPosts[0].Movie.Title}, {Key: "poster_path", Value: expectedPosts[0].Movie.PosterPath}, {Key: "tagline", Value: expectedPosts[0].Movie.Tagline}}},
			{Key: "content", Value: expectedPosts[0].Content},
		})
		second := mtest.CreateCursorResponse(1, ns, mtest.NextBatch, bson.D{
			{Key: "_id", Value: expectedPosts[1].ID},
			{Key: "user", Value: defaultUser},
			{Key: "movie", Value: bson.D{{Key: "id", Value: expectedPosts[1].Movie.ID}, {Key: "title", Value: expectedPosts[1].Movie.Title}, {Key: "poster_path", Value: expectedPosts[1].Movie.PosterPath}, {Key: "tagline", Value: expectedPosts[1].Movie.Tagline}}},
			{Key: "content", Value: expectedPosts[1].Content},
		})
		third := mtest.CreateCursorResponse(0, ns, mtest.NextBatch, bson.D{
			{Key: "_id", Value: expectedPosts[2].ID},
			{Key: "user", Value: defaultUser},
			{Key: "movie", Value: bson.D{{Key: "id", Value: expectedPosts[2].Movie.ID}, {Key: "title", Value: expectedPosts[2].Movie.Title}, {Key: "poster_path", Value: expectedPosts[2].Movie.PosterPath}, {Key: "tagline", Value: expectedPosts[2].Movie.Tagline}}},
			{Key: "content", Value: expectedPosts[2].Content},
		})
		mt.AddMockResponses(first, second, third)

		// Create a PostStore instance
		store := &PostStore{Collection: mt.Coll}

		// Call the function that we are testing
		posts, err := store.FindAllPosts(context.TODO(), 3)

		// Assert the function did not return an error
		assert.Nil(t, err)

		// Assert the function returned the expected posts
		assert.Equal(t, expectedPosts, posts)
	})
}

func TestFindPost(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("find post by ID", func(mt *mtest.T) {
		// Mock the expected result returned from the FindOne() function
		ns := mt.DB.Name() + "." + mt.Coll.Name()
		first := mtest.CreateCursorResponse(0, ns, mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedPosts[0].ID},
			{Key: "user", Value: defaultUser},
			{Key: "movie", Value: bson.D{{Key: "id", Value: expectedPosts[0].Movie.ID}, {Key: "title", Value: expectedPosts[0].Movie.Title}, {Key: "poster_path", Value: expectedPosts[0].Movie.PosterPath}, {Key: "tagline", Value: expectedPosts[0].Movie.Tagline}}},
			{Key: "content", Value: expectedPosts[0].Content},
		})
		mt.AddMockResponses(first)

		// Create a PostStore instance
		store := &PostStore{Collection: mt.Coll}

		// Call the function that we are testing
		post, err := store.FindPostByID(expectedPosts[0].ID.Hex(), context.TODO())

		// Assert the function did not return an error
		assert.Nil(t, err)

		// Assert the function returned the expected result
		assert.Equal(t, expectedPosts[0], post)
	})
}

func TestUpdateUserPost(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("update user post", func(mt *mtest.T) {
		// Mock the expected result returned from the FindOneAndUpdate() function
		expectedPost := bson.D{
			{Key: "_id", Value: expectedPosts[0].ID},
			{Key: "user", Value: defaultUser},
			{Key: "movie", Value: bson.D{{Key: "id", Value: expectedPosts[0].Movie.ID}, {Key: "title", Value: expectedPosts[0].Movie.Title}, {Key: "poster_path", Value: expectedPosts[0].Movie.PosterPath}, {Key: "tagline", Value: expectedPosts[0].Movie.Tagline}}},
			{Key: "content", Value: expectedPosts[0].Content},
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "value", Value: expectedPost}))

		// Create a PostStore instance
		store := &PostStore{Collection: mt.Coll}

		// Call the function that we are testing
		err := store.UpdateUserPost(expectedPosts[0].ID.Hex(), expectedPosts[0].Content, expectedPosts[0].User.ID, context.TODO())

		// Assert the function did not return an error
		assert.Nil(t, err)

		// Assert the post content was updated
		assert.Equal(t, expectedPosts[0].Content, "Test Content")
	})
}

func TestDeleteUserPost(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		// Mock the expected result returned from the DeleteOne() function
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "acknowledged", Value: true}, bson.E{Key: "n", Value: 1}))

		// Create a PostStore instance
		store := &PostStore{Collection: mt.Coll}

		// Call the function that we are testing
		err := store.DeleteUserPost(expectedPosts[0].ID.Hex(), expectedPosts[0].User, context.TODO())

		// Assert the function did not return an error
		assert.Nil(t, err)
	})

	mt.Run("no posts deleted", func(mt *mtest.T) {
		// Mock the expected result returned from the DeleteOne() function
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "acknowledged", Value: true}, bson.E{Key: "n", Value: 0}))

		// Create a PostStore instance
		store := &PostStore{Collection: mt.Coll}

		// Call the function that we are testing
		err := store.DeleteUserPost(expectedPosts[0].ID.Hex(), expectedPosts[0].User, context.TODO())

		// Assert the function returned an error
		assert.NotNil(t, err)
	})
}

func TestSave(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		// Mock the expected result returned from the InsertOne() function
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Create a PostStore instance
		store := &PostStore{Collection: mt.Coll}

		// Call the function that we are testing
		err := store.Save(context.TODO(), expectedPosts[0])

		// Assert the function did not return an error
		assert.Nil(t, err)
	})
}
