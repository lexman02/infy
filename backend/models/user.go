package models

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"infy/db"
	"os"
	"time"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	IsAdmin  bool               `json:"isAdmin" bson:"isAdmin"`
}

// NewUser creates a new user instance
func NewUser(username, email, password string) *User {
	return &User{ID: primitive.NewObjectID(), Username: username, Email: email, Password: password}
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
