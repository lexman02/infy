package controllers

import (
	"infy/api"
	"infy/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPosts(c *gin.Context) {
	// Get all posts
	posts, err := models.FindAllPosts(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Create a response with the post and the created date
	var postsResponse []map[string]interface{}
	for _, post := range posts {
		postsResponse = append(postsResponse, map[string]interface{}{
			"post":    post,
			"created": post.ID.Timestamp().Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(200, postsResponse)
}

func GetPost(c *gin.Context) {
	// Get the post by ID
	post, err := models.FindPostByID(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}
	// Get all comments for the post since it's the post details
	comments, err := models.FindCommentsByPostID(post.ID.Hex(), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	postResponse := map[string]interface{}{
		"post":     post,
		"created":  post.ID.Timestamp().Format("2006-01-02 15:04:05"),
		"comments": comments,
	}

	c.JSON(200, postResponse)
}

func CreatePost(c *gin.Context) {
	var post struct {
		MovieID string `json:"movie_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// Bind the request body to the post struct
	if err := c.ShouldBindJSON(&post); err != nil {
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

	// Get the movie details
	movie, err := api.GetMovieDetails(post.MovieID)
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Create the post
	newPost := models.NewPost(user.(*models.User), movie, post.Content)
	if err := newPost.Save(c.Request.Context()); err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	c.JSON(200, newPost)
}

func UpdatePost(c *gin.Context) {
	var post struct {
		Content string `json:"content" binding:"required"`
	}

	// Bind the request body to the post struct
	if err := c.ShouldBindJSON(&post); err != nil {
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

	// Update the post
	err := models.UpdateUserPost(c.Param("id"), post.Content, user.(*models.User).ID, c.Request.Context())
	if err != nil {
		// Check if the post was not found or the user is not the author
		if err == mongo.ErrNoDocuments {
			c.JSON(403, gin.H{"error": "Post not found or you are not the author of this post"})
			return
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

	c.JSON(200, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	// Delete the post
	err := models.DeleteUserPost(c.Param("id"), user.(*models.User).ID, c.Request.Context())
	if err != nil {
		// Check if the post was not found or the user is not the author
		if err == mongo.ErrNoDocuments {
			c.JSON(403, gin.H{"error": "Post not found or you are not the author of this post"})
			return
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

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
