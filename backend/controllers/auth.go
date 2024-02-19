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

func Login(c *gin.Context) {
	var login struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	user, err := models.FindUserByEmail(login.Email, c.Request.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	expTime := time.Now().Add(24 * time.Hour)
	token, err := user.GetJwtToken(expTime)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		log.Println(err)
		return
	}

	c.SetCookie("token", token, int(expTime.Unix()), "/", "localhost", utils.IsProd(), true)
	c.JSON(200, gin.H{"success": "Logged in"})
}

func Signup(c *gin.Context) {
	var signup struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&signup); err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	user, _ := models.FindUserByEmail(signup.Email, c.Request.Context())
	if user != nil {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	user = models.NewUser(signup.Username, signup.Email, string(hashedPassword))
	err = user.Save(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User created"})
}

func Home(c *gin.Context) {

}

func Premium(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
