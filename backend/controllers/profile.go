package controllers

import (
	"infy/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"profile": user.Profile})
}

func GetProfile(c *gin.Context) {
	profile, err := models.FindUserProfileByUsername(c.Param("username"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"profile": profile})
}

func Follow(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := user.FollowUser(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User followed"})
}

func Unfollow(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := user.UnfollowUser(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"success": "User unfollowed"})
}
