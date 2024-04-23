package controllers

import (
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetReportedPosts fetches a list of reported posts up to a specified limit.
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

// DeleteReportedPost allows an admin to delete a post identified by its ID.
func DeleteReportedPost(c *gin.Context) {
	postID := c.Param("id")

	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "User not found"})
		return
	}

	err := models.DeleteUserPost(postID, user.(*models.User), c.Request.Context())
	if err != nil {
		// Handle specific errors based on the MongoDB response.
		if err == mongo.ErrNoDocuments {
			// Attempt to remove the post as it was reported.
			err := models.RemoveReportedPost(postID, c.Request.Context())
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to delete post"})
				log.Println(err)
				return
			}

			c.JSON(200, gin.H{"message": "Post deleted successfully"})
			return
		}

		// Check if the post ID is invalid (not a valid MongoDB Hex ID).
		if err == primitive.ErrInvalidHex {
			c.JSON(400, gin.H{"error": "Invalid post ID"})
			return
		}

		c.JSON(500, gin.H{"error": "Failed to delete post"})
		log.Println(err)
		return
	}

	// Successfully remove the reported post.
	err = models.RemoveReportedPost(postID, c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

// ToggleAdminStatus allows a superadmin to toggle the admin status of a user.
func ToggleAdminStatus(c *gin.Context) {
	userID := c.Param("id")

	err := models.ToggleAdmin(userID, c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update admin status"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin status updated successfully"})
}

// GetUsers retrieves all users from the database.
func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
