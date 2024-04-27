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

// GetPosts retrieves a list of posts and enriches them with user-specific reaction data.
func GetPosts(c *gin.Context) {
	// Retrieve posts with a default limit for pagination
	posts, err := models.FindAllPosts(c.Request.Context(), 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts"})
		log.Println(err)
		return
	}

	// Attempt to extract the JWT token from cookies
	token, err := c.Cookie("token")
	if err != nil {
		token = ""
	}

	var userID primitive.ObjectID
	// Decode token to fetch user ID if token is present
	if token != "" {
		user, err := middleware.GetUserFromToken(token, c)
		if err != nil {
			if err.Error() == "token deleted due to expiration" {
				return
			}
			c.JSON(500, gin.H{"error": "An error occurred"})
			log.Println(err)
			return
		}
		userID = user.ID
	}

	// Prepare posts for JSON response
	var postsResponse []map[string]interface{}
	for _, post := range posts {
		var likes, dislikes int
		var liked, disliked bool

		// Count likes and dislikes for each post
		for _, reaction := range post.Reactions {
			if reaction.Liked {
				likes++
			}
			if reaction.Disliked {
				dislikes++
			}
			// Check if the current user has liked or disliked the post
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

// GetPost retrieves a single post by ID and its associated comments, including user-specific reaction data.
func GetPost(c *gin.Context) {
	// Get the post by ID
	limit := int64(20)
	post, err := models.FindPostByID(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
		log.Println(err)
		return
	}
	// Get all comments for the post since it's the post details
	comments, err := models.FindCommentsByPostID(post.ID.Hex(), c.Request.Context(), limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "An error occurred"})
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
		user, err := middleware.GetUserFromToken(token, c)
		if err != nil {
			if err.Error() == "token deleted due to expiration" {
				return
			}
			c.JSON(500, gin.H{"error": "An error occurred"})
			log.Println(err)
			return
		}

		userID = user.ID
	}

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

	postResponse := map[string]interface{}{
		"post":     post,
		"liked":    liked,
		"disliked": disliked,
		"likes":    likes,
		"dislikes": dislikes,
		"created":  post.ID.Timestamp().Format("2006-01-02 15:04:05"),
		"comments": comments,
	}

	c.JSON(200, postResponse)
}

// GetPostsByMovieID fetches all posts related to a specific movie by the movie's ID.
func GetPostsByMovieID(c *gin.Context) {
	movieID := c.Param("movieID") // Extracting movieID from the URL parameter
	limit := int64(20)

	posts, err := models.FindPostsByMovieID(movieID, c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve posts for the movie"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// CreatePost handles the creation of a new post related to a movie.
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
	_, movie, err := api.GetMovieDetails(post.MovieID, true)
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

// UpdatePost allows authorized users to modify an existing post.
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

// DeletePost allows authorized users to delete an existing post.
func DeletePost(c *gin.Context) {
	// Get the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "An error occurred"})
		return
	}

	// Delete the post
	err := models.DeleteUserPost(c.Param("id"), user.(*models.User), c.Request.Context())
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

// LikePost handles the action of a user liking a post.
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

// DislikePost handles the action of a user disliking a post.
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

// GetUserPosts retrieves posts created by a specific user.
func GetUserPosts(c *gin.Context) {
	userID := c.Param("userID")
	limit := int64(20)

	posts, err := models.FindPostsByUserID(userID, c.Request.Context(), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user's posts"})
		return
	}

	// Get the token from the cookie
	token, err := c.Cookie("token")
	if err != nil {
		token = ""
	}

	// Get the user ID from the token
	var authenticatedUserID primitive.ObjectID
	if token != "" {
		user, err := middleware.GetUserFromToken(token, c)
		if err != nil {
			if err.Error() == "token deleted due to expiration" {
				return
			}
			c.JSON(500, gin.H{"error": "An error occurred"})
			log.Println(err)
			return
		}

		authenticatedUserID = user.ID
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

			if reaction.UserID == authenticatedUserID {
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

	c.JSON(http.StatusOK, postsResponse)
}

// ReportPost allows a user to report a post as inappropriate.
func ReportPost(c *gin.Context) {
	// Attempt to report the post and handle errors
	err := models.ReportPost(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to report post"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post reported successfully"})
}
