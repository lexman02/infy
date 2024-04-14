package controllers

import (
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetReportedPosts(c *gin.Context) {
	limit := int64(20)

	reportedPosts, err := models.FindReportedPosts(c.Request.Context(), limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, reportedPosts)
}

// DeleteReportedPost allows admin to delete a post
func DeleteReportedPost(c *gin.Context) {
	postID := c.Param("id")

	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	err := models.DeleteUserPost(postID, user.(*models.User), c.Request.Context())
	if err != nil {
		// Check if the post was not found
		if err == mongo.ErrNoDocuments {
			// Remove the reported post
			err := models.RemoveReportedPost(postID, c.Request.Context())
			if err != nil {
				c.JSON(500, gin.H{"error": "An error occurred"})
				log.Println(err)
				return
			}

			c.JSON(200, gin.H{"message": "Post deleted successfully"})
		}

		// Check if the post ID is invalid
		if err == primitive.ErrInvalidHex {
			c.JSON(400, gin.H{"error": "Invalid post ID"})
			return
		}

		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Remove the reported post
	err = models.RemoveReportedPost(postID, c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

// ToggleAdminStatus allows superadmin to toggle admin status of a user
func ToggleAdminStatus(c *gin.Context) {
	userID := c.Param("id")

	err := models.ToggleAdmin(userID, c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin status updated successfully"})
}

// Get all users
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
