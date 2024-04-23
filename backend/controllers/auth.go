package controllers

import (
	"errors"
	"infy/models"
	"infy/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Login authenticates a user by checking credentials and sets a JWT token as a cookie if successful.
func Login(c *gin.Context) {
	var login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the request body to the login struct
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(500, gin.H{"error": "Invalid request data"})
		log.Println(err)
		return
	}

	// Retrieve the user by email
	user, err := models.FindUserByEmail(login.Email, c.Request.Context())
	if err != nil {
		// Handle not found error specifically
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token and set it as a cookie
	expTime := time.Now().Add(24 * time.Hour)
	token, err := user.GetJwtToken(expTime)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		log.Println(err)
		return
	}

	c.SetCookie("token", token, int(expTime.Unix()), "/", utils.GetEnv("SITE_DOMAIN", "localhost"), utils.IsProd(), true)
	c.JSON(200, gin.H{"success": "Logged in"})
}

// Signup creates a new user account with the provided details after verifying that the user does not already exist.
func Signup(c *gin.Context) {
	var signup struct {
		Email           string `json:"email" binding:"required"`
		Username        string `json:"username" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
		FirstName       string `json:"first_name" binding:"required"`
		LastName        string `json:"last_name" binding:"required"`
		DateOfBirth     string `json:"date_of_birth" binding:"required"`
	}

	// Bind and validate request body
	if err := c.ShouldBindJSON(&signup); err != nil {
		c.JSON(500, gin.H{"error": "Invalid request data"})
		log.Println(err)
		return
	}

	// Ensure passwords match
	if signup.Password != signup.ConfirmPassword {
		c.JSON(400, gin.H{"error": "Passwords do not match"})
		return
	}

	// Check if the user already exists
	user, _ := models.FindUserByEmail(signup.Email, c.Request.Context())
	if user != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		log.Println(err)
		return
	}

	// Parse the date of birth and create the user profile
	dateOfBirth, err := time.Parse("2006-01-02", signup.DateOfBirth)
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid date format"})
		log.Println(err)
		return
	}

	profile := models.NewProfile(signup.FirstName, signup.LastName, dateOfBirth, models.NewPreferences())
	user = models.NewUser(signup.Username, signup.Email, string(hashedPassword), profile)

	// Save the new user and handle potential errors
	err = user.Save(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		log.Println(err)
		return
	}

	// Generate a JWT token and set it as a cookie
	expTime := time.Now().Add(24 * time.Hour)
	token, err := user.GetJwtToken(expTime)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		log.Println(err)
		return
	}

	c.SetCookie("token", token, int(expTime.Unix()), "/", utils.GetEnv("SITE_DOMAIN", "localhost"), utils.IsProd(), true)
	c.JSON(200, gin.H{"success": "User created"})
}

// User retrieves and displays the current authenticated user's details.
func User(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

// Logout terminates the user session by clearing the authentication token cookie.
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", utils.GetEnv("SITE_DOMAIN", "localhost"), utils.IsProd(), true)
	c.JSON(200, gin.H{"success": "Logged out"})
}
