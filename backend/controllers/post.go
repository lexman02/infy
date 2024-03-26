package controllers

import (
	"infy/api"
	"infy/middleware"
	"infy/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPosts(c *gin.Context) {
	// Get the posts from the database
	posts, err := models.FindAllPosts(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	// Get the token from the cookie
	token, err := c.Cookie("token")
	if err != nil {
		token = ""
	}

	// Get the user ID from the token
	var userID primitive.ObjectID
	if token != "" {
		middleware.Authorized()(c)
		user, exists := c.Get("user")
		if exists {
			userID = user.(*models.User).ID
		}
	}

	// Create a response with the post and the created date
	var postsResponse []map[string]interface{}
	for _, post := range posts {
		// Like and dislike counters
		var likes, dislikes int = 0, 0
		// Like and dislike status
		var liked, disliked bool = false, false

		// Check if the user has liked the post and increment the like or dislike counters
		for _, reaction := range post.Reactions {
			if reaction.Liked {
				likes++
			}

			if reaction.Disliked {
				dislikes++
			}

			if reaction.UserID == userID {
				liked = reaction.Liked
				disliked = reaction.Disliked
			}
		}

		// Append the post to the response
		postsResponse = append(postsResponse, map[string]interface{}{
			"post":     post,
			"liked":    liked,
			"disliked": disliked,
			"likes":    likes,
			"dislikes": dislikes,
			"created":  post.ID.Timestamp().Format("2006-01-02 15:04:05"),
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

func LikePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println("User not found in context")
		return
	}

	var reaction struct {
		IsLiked bool `json:"is_liked"`
	}

	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(400, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	err := models.UpdateReaction(c.Param("id"), user.(*models.User).ID, !reaction.IsLiked, false, c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}
}

func DislikePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println("User not found in context")
		return
	}

	var reaction struct {
		IsDisliked bool `json:"is_disliked"`
	}

	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(400, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}

	err := models.UpdateReaction(c.Param("id"), user.(*models.User).ID, false, !reaction.IsDisliked, c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}
}

func GetUserPosts(c *gin.Context) {
	userID := c.Param("userID") // Extracting userID from the URL parameter

	posts, err := models.FindPostsByUserID(userID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user's posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
