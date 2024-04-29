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
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"-" bson:"password"`
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
	WatchList []string             `json:"watchlist"`                            // Movie IDs
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

// NewPreferences creates a new preferences instance with empty slices
func NewPreferences() *Preferences {
	return &Preferences{Genres: []string{}, Following: nil, Followers: nil, WatchList: []string{}, Watched: []string{}}
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

func FindUserByUsername(username string, ctx context.Context) (*User, error) {
	var user User

	err := db.UsersCollection().FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FollowUser adds a user to the following list and updates the other user's followers list
func (u *User) FollowUser(id string, ctx context.Context) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	followUser := bson.M{"$addToSet": bson.M{"profile.preferences.following": userId}} // Prevent duplicates
	_, err = db.UsersCollection().UpdateByID(ctx, u.ID, followUser)
	if err != nil {
		return err
	}

	followerUser := bson.M{"$addToSet": bson.M{"profile.preferences.followers": u.ID}} // Prevent duplicates
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

	unfollowUser := bson.M{"$pull": bson.M{"profile.preferences.following": userId}}
	_, err = db.UsersCollection().UpdateByID(ctx, u.ID, unfollowUser)
	if err != nil {
		return err
	}

	unfollowerUser := bson.M{"$pull": bson.M{"profile.preferences.followers": u.ID}}
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
func AddMovieToWatchedList(userID, movieID string, ctx context.Context) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error converting userID to ObjectID: %v", err)
		return err
	}

	update := bson.M{"$addToSet": bson.M{"profile.preferences.watched": movieID}} // Prevent duplicates
	_, err = db.UsersCollection().UpdateByID(ctx, userObjectID, update)

	return err
}

// AddMovieToWatchlist adds a movie ID to the user's watchlist
func AddMovieToWatchlist(userID, movieID string, ctx context.Context) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$addToSet": bson.M{"profile.preferences.watchlist": movieID}} // Prevent duplicates
	_, err = db.UsersCollection().UpdateByID(ctx, userObjectID, update)

	return err
}

// RemoveMovieFromWatchedList removes a movie ID from the user's watched list
func RemoveMovieFromWatchedList(userID, movieID string, ctx context.Context) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$pull": bson.M{"profile.preferences.watched": movieID}}
	_, err = db.UsersCollection().UpdateByID(ctx, userObjectID, update)

	return err
}

// RemoveMovieFromWatchlist removes a movie ID from the user's watchlist
func RemoveMovieFromWatchlist(userID, movieID string, ctx context.Context) error {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	update := bson.M{"$pull": bson.M{"profile.preferences.watchlist": movieID}}
	_, err = db.UsersCollection().UpdateByID(ctx, userObjectID, update)

	return err
}

func FindFollowedWhoWatchedMovie(userID, movieID string) ([]User, error) {
	var users []User
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$match", Value: bson.M{"_id": userObjectID}}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from": "users",
			"let":  bson.M{"followingIds": "$preferences.following"},
			"pipeline": mongo.Pipeline{
				bson.D{{Key: "$match", Value: bson.M{"$expr": bson.M{"$in": []interface{}{"$_id", "$$followingIds"}}}}},
				bson.D{{Key: "$match", Value: bson.M{"preferences.watched": movieID}}},
			},
			"as": "followedWhoWatched",
		}}},
		bson.D{{Key: "$unwind", Value: "$followedWhoWatched"}},
		bson.D{{Key: "$replaceRoot", Value: bson.M{"newRoot": "$followedWhoWatched"}}},
	}

	ctx := context.TODO()
	cursor, err := db.UsersCollection().Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// ToggleAdmin toggles the admin status of a user
func ToggleAdmin(userID string, ctx context.Context) error {
	// Convert the user ID to an ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// Check if the user is already an admin
	user, err := FindUserByID(userID, ctx)
	if err != nil {
		return err
	}

	// Update to the opposite of the current value
	update := bson.M{"$set": bson.M{"isAdmin": !user.IsAdmin}}
	_, err = db.UsersCollection().UpdateOne(ctx, bson.M{"_id": userObjectID}, update)
	if err != nil {
		return err
	}

	return nil
}

// GetUsers returns all users
func GetUsers() ([]User, error) {
	var users []User
	cursor, err := db.UsersCollection().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

// AddAvatar adds an avatar to the user's profile
func AddAvatar(user *User, filepath string, ctx context.Context) error {
	update := bson.M{"$set": bson.M{"profile.avatar": filepath}}
	_, err := db.UsersCollection().UpdateByID(ctx, user.ID, update)
	if err != nil {
		return err
	}

	err = refreshUser(user, ctx)
	if err != nil {
		return err
	}

	return nil
}

// refreshUser updates the user in all posts and comments
func refreshUser(user *User, ctx context.Context) error {
	// Update the user in all posts
	update := bson.M{"$set": bson.M{"user": user}}
	_, err := db.PostsCollection().UpdateMany(ctx, bson.M{"user._id": user.ID}, update)
	if err != nil {
		return err
	}

	// Update the user in all comments
	update = bson.M{"$set": bson.M{"user": user}}
	_, err = db.CommentsCollection().UpdateMany(ctx, bson.M{"user._id": user.ID}, update)
	if err != nil {
		return err
	}

	return nil
}
