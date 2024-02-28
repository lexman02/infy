package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"infy/models"
	"infy/utils"
	"log"
	"time"
)

// Login checks if the user exists and compares the password with the hashed password and sets a JWT token as a cookie
func Login(c *gin.Context) {
	var login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the request body to the login struct
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Find the user by email
	user, err := models.FindUserByEmail(login.Email, c.Request.Context())
	if err != nil {
		// If the user is not found, return an error
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Compare the password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	expTime := time.Now().Add(24 * time.Hour)
	token, err := user.GetJwtToken(expTime)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		log.Println(err)
		return
	}

	// Set the token as a cookie
	c.SetCookie("token", token, int(expTime.Unix()), "/", utils.GetEnv("SITE_DOMAIN", "localhost"), utils.IsProd(), true)
	c.JSON(200, gin.H{"success": "Logged in"})
}

// Signup checks if the user already exists and creates a new user
func Signup(c *gin.Context) {
	var signup struct {
		Email           string `json:"email" binding:"required"`
		Username        string `json:"username" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	// Bind the request body to the signup struct
	if err := c.ShouldBindJSON(&signup); err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Check if the passwords match
	if signup.Password != signup.ConfirmPassword {
		c.JSON(500, gin.H{"error": "Passwords do not match"})
		return
	}

	// Check if the user already exists by email
	user, _ := models.FindUserByEmail(signup.Email, c.Request.Context())
	if user != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Create the user
	user = models.NewUser(signup.Username, signup.Email, string(hashedPassword))
	err = user.Save(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User created"})
}

func User(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	type userResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		IsAdmin  bool   `json:"isAdmin"`
	}

	c.JSON(200, gin.H{"user": userResponse{
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
	}})
}

func Logout(c *gin.Context) {

}
