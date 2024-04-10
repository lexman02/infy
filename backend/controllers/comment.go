package controllers

import (
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateComment(c *gin.Context) {
	var comment struct {
		PostID  string `json:"post_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// Bind the request body to the post struct
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	// Parse the post ID
	postID, err := primitive.ObjectIDFromHex(comment.PostID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}

	// Create the comment
	newComment := models.NewComment(postID, user.(*models.User), comment.Content)
	if err := newComment.Save(c.Request.Context()); err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, newComment)
}

func UpdateComment(c *gin.Context) {
	var comment struct {
		Content string `json:"content" binding:"required"`
	}

	// Bind the request body to the comment struct
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	// Update the comment
	err := models.UpdateUserComment(c.Param("id"), comment.Content, user.(*models.User).ID, c.Request.Context())
	if err != nil {
		// Check if the post was not found or the user is not the author
		if err == mongo.ErrNoDocuments {
			c.JSON(403, gin.H{"error": "Comment not found or you are not the author of this post"})
			return
		}

		// Check if the comment ID is invalid
		if err == primitive.ErrInvalidHex {
			c.JSON(400, gin.H{"error": "Invalid comment ID"})
			return
		}

		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Comment updated successfully"})
}

func DeleteComment(c *gin.Context) {
	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	// Delete the comment
	err := models.DeleteUserComment(c.Param("id"), user.(*models.User).ID, c.Request.Context())
	if err != nil {
		// Check if the comment was not found or the user is not the author
		if err == mongo.ErrNoDocuments {
			c.JSON(403, gin.H{"error": "Comment not found or you are not the author of this post"})
			return
		}

		// Check if the comment ID is invalid
		if err == primitive.ErrInvalidHex {
			c.JSON(400, gin.H{"error": "Invalid comment ID"})
			return
		}

		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}

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
