package controllers

import (
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateComment processes the incoming request to create a new comment on a post.
func CreateComment(c *gin.Context) {
	var comment struct {
		PostID  string `json:"post_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// Bind JSON payload to the struct and handle binding errors
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		log.Println(err)
		return
	}

	// Retrieve authenticated user from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "User context not found"})
		return
	}

	// Validate post ID format
	postID, err := primitive.ObjectIDFromHex(comment.PostID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}

	// Instantiate and save new comment
	newComment := models.NewComment(postID, user.(*models.User), comment.Content)
	if err := newComment.Save(c.Request.Context()); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save comment"})
		log.Println(err)
		return
	}

	c.JSON(200, newComment)
}

// UpdateComment modifies an existing comment based on the comment ID provided in the URL.
func UpdateComment(c *gin.Context) {
	var comment struct {
		Content string `json:"content" binding:"required"`
	}

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		log.Println(err)
		return
	}

	// Retrieve authenticated user from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "User context not found"})
		return
	}

	// Attempt to update the comment
	err := models.UpdateUserComment(c.Param("id"), comment.Content, user.(*models.User).ID, c.Request.Context())
	if err != nil {
		// Handle no documents and invalid ID errors specifically
		if err == mongo.ErrNoDocuments {
			c.JSON(403, gin.H{"error": "Comment not found or unauthorized modification attempt"})
			return
		}
		if err == primitive.ErrInvalidHex {
			c.JSON(400, gin.H{"error": "Invalid comment ID"})
			return
		}

		c.JSON(500, gin.H{"error": "Failed to update comment"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Comment updated successfully"})
}

// DeleteComment removes a comment based on the comment ID provided in the URL.
func DeleteComment(c *gin.Context) {
	// Retrieve authenticated user from context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "User context not found"})
		return
	}

	// Attempt to delete the comment
	err := models.DeleteUserComment(c.Param("id"), user.(*models.User).ID, c.Request.Context())
	if err != nil {
		// Handle no documents and invalid ID errors specifically
		if err == mongo.ErrNoDocuments {
			c.JSON(403, gin.H{"error": "Comment not found or unauthorized deletion attempt"})
			return
		}
		if err == primitive.ErrInvalidHex {
			c.JSON(400, gin.H{"error": "Invalid comment ID"})
			return
		}

		c.JSON(500, gin.H{"error": "Failed to delete comment"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}

// LikeComment toggles the 'like' status of a comment for the authenticated user.
func LikeComment(c *gin.Context) {
	commentID := c.Param("commentId")
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := models.ToggleLikeOnComment(commentID, userID.Hex(), true); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment liked successfully"})
}

// DislikeComment toggles the 'dislike' status of a comment for the authenticated user.
func DislikeComment(c *gin.Context) {
	commentID := c.Param("commentId")
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := models.ToggleLikeOnComment(commentID, userID.Hex(), false); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dislike comment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment disliked successfully"})
}

// getUserIDFromContext extracts the user ID from the authenticated user context.
func getUserIDFromContext(c *gin.Context) (primitive.ObjectID, bool) {
	user, exists := c.Get("user")
	if !exists {
		return primitive.NilObjectID, false
	}
	if userModel, ok := user.(*models.User); ok {
		return userModel.ID, true
	}
	return primitive.NilObjectID, false
}
