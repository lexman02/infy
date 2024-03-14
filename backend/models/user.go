package models

import (
	"context"
	"infy/db"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"-" bson:"-"`
	IsAdmin  bool               `json:"isAdmin" bson:"isAdmin"`
	Profile  Profile            `json:"profile" bson:"profile"`
}

type Profile struct {
	FirstName   string      `json:"first_name" bson:"first_name"`
	LastName    string      `json:"last_name" bson:"last_name"`
	DateOfBirth time.Time   `json:"date_of_birth" bson:"date_of_birth"`
	Avatar      string      `json:"avatar" bson:"avatar,omitempty"`
	Rank        string      `json:"rank" bson:"rank"`
	Preferences Preferences `json:"preferences" bson:"preferences"`
}

type Preferences struct {
	Genres    []string             `json:"genres"`
	Following []primitive.ObjectID `json:"following" bson:"following,omitempty"` // IDs of users
	Followers []primitive.ObjectID `json:"followers" bson:"followers,omitempty"` // IDs of users
	WatchList []string             `json:"watch_list"`                           // Movie IDs
	Watched   []string             `json:"watched"`                              // Movie IDs
}

// NewUser creates a new user instance
func NewUser(username, email, password string, profile *Profile) *User {
	return &User{ID: primitive.NewObjectID(), Username: username, Email: email, Password: password, IsAdmin: false, Profile: *profile}
}

// NewProfile creates a new profile instance
func NewProfile(firstName, lastName string, dateOfBirth time.Time, preferences *Preferences) *Profile {
	return &Profile{FirstName: firstName, LastName: lastName, DateOfBirth: dateOfBirth, Avatar: "", Rank: "Newbie", Preferences: *preferences}
}

//// NewPreferences creates a new preferences instance
//func NewPreferences(genres, watchList, watched []string) *Preferences {
//	return &Preferences{Genres: genres, Following: nil, Followers: nil, WatchList: watchList, Watched: watched}
//}

// NewPreferences creates a new preferences instance with empty slices
func NewPreferences() *Preferences {
	return &Preferences{Genres: nil, Following: nil, Followers: nil, WatchList: nil, Watched: nil}
}

// FindUserByEmail finds a user by email
func FindUserByEmail(email string, ctx context.Context) (*User, error) {
	var user User

	err := db.UsersCollection().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByID finds a user by ID
func FindUserByID(id string, ctx context.Context) (*User, error) {
	var user User

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = db.UsersCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)

	return &user, err
}

func FindUserProfileByUsername(username string, ctx context.Context) (*Profile, error) {
	var user User

	err := db.UsersCollection().FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user.Profile, nil
}

// FollowUser adds a user to the following list and updates the other user's followers list
func (u *User) FollowUser(id string, ctx context.Context) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	followUser := bson.M{"$addToSet": bson.M{"following": userId}} // Prevent duplicates
	_, err = db.UsersCollection().UpdateByID(ctx, u.ID, followUser)
	if err != nil {
		return err
	}

	followerUser := bson.M{"$addToSet": bson.M{"followers": u.ID}} // Prevent duplicates
	_, err = db.UsersCollection().UpdateByID(ctx, userId, followerUser)
	if err != nil {
		return err
	}

	return nil
}

// UnfollowUser removes a user from the following list and updates the other user's followers list
func (u *User) UnfollowUser(id string, ctx context.Context) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	unfollowUser := bson.M{"$pull": bson.M{"following": userId}}
	_, err = db.UsersCollection().UpdateByID(ctx, u.ID, unfollowUser)
	if err != nil {
		return err
	}

	unfollowerUser := bson.M{"$pull": bson.M{"followers": u.ID}}
	_, err = db.UsersCollection().UpdateByID(ctx, userId, unfollowerUser)
	if err != nil {
		return err
	}

	return nil
}

// GetJwtToken returns a JWT token with the user's ID as the subject
func (u *User) GetJwtToken(exp time.Time) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   u.ID.Hex(),
		ExpiresAt: jwt.NewNumericDate(exp),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Save saves a user to the database
func (u *User) Save(ctx context.Context) error {
	_, err := db.UsersCollection().InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

// AddMovieToWatchedList adds a movie ID to the user's watched list
func AddMovieToWatchedList(userID, movieID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error converting userID to ObjectID: %v", err)
		return err
	}

	update := bson.M{"$addToSet": bson.M{"watched": movieID}} // Prevent duplicates
	result, err := db.UsersCollection().UpdateByID(context.Background(), userObjectID, update)
	if err != nil {
		log.Printf("Error adding movie to watched list: %v", err)
		return err
	}

	// Now 'result' is defined, and you can log its properties
	log.Printf("Updated document count: %v", result.ModifiedCount)
	return nil
}

// AddMovieToWatchlist adds a movie ID to the user's watchlist
func AddMovieToWatchlist(userID, movieID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$addToSet": bson.M{"watchlist": movieID}} // Prevent duplicates
	_, err = db.UsersCollection().UpdateByID(context.Background(), userObjectID, update)

	return err
}

// RemoveMovieFromWatchedList removes a movie ID from the user's watched list
func RemoveMovieFromWatchedList(userID, movieID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$pull": bson.M{"watched": movieID}}
	_, err = db.UsersCollection().UpdateByID(context.Background(), userObjectID, update)

	return err
}

// RemoveMovieFromWatchlist removes a movie ID from the user's watchlist
func RemoveMovieFromWatchlist(userID, movieID string) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$pull": bson.M{"watchlist": movieID}}
	_, err = db.UsersCollection().UpdateByID(context.Background(), userObjectID, update)

	return err
}
