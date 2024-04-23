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
		user, err := middleware.GetUserFromToken(token, c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode token"})
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

		// Append structured post data to response slice
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

// GetPost retrieves a single post by ID and its associated comments, including user-specific reaction data.
func GetPost(c *gin.Context) {
	// Retrieve the post by its ID
	post, err := models.FindPostByID(c.Param("id"), c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Post not found"})
		log.Println(err)
		return
	}

	// Fetch comments related to the post
	comments, err := models.FindCommentsByPostID(post.ID.Hex(), c.Request.Context(), 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		token = ""
	}

	var userID primitive.ObjectID
	// Decode token to fetch user ID if token is present
	if token != "" {
		user, err := middleware.GetUserFromToken(token, c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode token"})
			log.Println(err)
			return
		}
		userID = user.ID
	}

	var likes, dislikes int
	var liked, disliked bool

	// Count likes and dislikes for the post and check if the current user has liked or disliked
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

	// Structure the full post data including comments for the response
	postResponse := map[string]interface{}{
		"post":     post,
		"liked":    liked,
		"disliked": disliked,
		"likes":    likes,
		"dislikes": dislikes,
		"created":  post.ID.Timestamp().Format("2006-01-02 15:04:05"),
		"comments": comments,
	}

	c.JSON(http.StatusOK, postResponse)
}

// GetPostsByMovieID fetches all posts related to a specific movie by the movie's ID.
func GetPostsByMovieID(c *gin.Context) {
	movieID := c.Param("movieID")
	// Fetch posts by movie ID with a default pagination limit
	posts, err := models.FindPostsByMovieID(movieID, c.Request.Context(), 20)
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

	// Bind incoming JSON to the struct and handle errors
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post data"})
		log.Println(err)
		return
	}

	// Fetch the user from the context, who is creating the post
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
		return
	}

	// Fetch movie details using the API
	_, movie, err := api.GetMovieDetails(post.MovieID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie details"})
		log.Println(err)
		return
	}

	// Create and save the new post
	newPost := models.NewPost(user.(*models.User), movie, post.Content)
	if err := newPost.Save(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the post"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, newPost)
}

// UpdatePost allows authorized users to modify an existing post.
func UpdatePost(c *gin.Context) {
	var post struct {
		Content string `json:"content" binding:"required"`
	}

	// Bind incoming JSON to the struct and handle errors
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		log.Println(err)
		return
	}

	// Fetch the user from the context, who is attempting the update
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
		return
	}

	// Attempt to update the post and handle errors, such as not finding the post or permission issues
	err := models.UpdateUserPost(c.Param("id"), post.Content, user.(*models.User).ID, c.Request.Context())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusForbidden, gin.H{"error": "No such post found or permission denied"})
			return
		}
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the post"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePost allows authorized users to delete an existing post.
func DeletePost(c *gin.Context) {
	// Fetch the user from the context, who is attempting the deletion
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
		return
	}

	// Attempt to delete the post and handle errors, such as not finding the post or permission issues
	err := models.DeleteUserPost(c.Param("id"), user.(*models.User), c.Request.Context())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusForbidden, gin.H{"error": "No such post found or permission denied"})
			return
		}
		if err == primitive.ErrInvalidHex {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the post"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// LikePost handles the action of a user liking a post.
func LikePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
		return
	}

	var reaction struct {
		IsLiked bool `json:"is_liked"`
	}

	// Bind incoming JSON to the struct and handle errors
	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction data"})
		return
	}

	// Update the post reaction to reflect the like and handle errors
	err := models.UpdateReaction(c.Param("id"), user.(*models.User).ID, reaction.IsLiked, false, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction updated successfully"})
}

// DislikePost handles the action of a user disliking a post.
func DislikePost(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
		return
	}

	var reaction struct {
		IsDisliked bool `json:"is_disliked"`
	}

	// Bind incoming JSON to the struct and handle errors
	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction data"})
		return
	}

	// Update the post reaction to reflect the dislike and handle errors
	err := models.UpdateReaction(c.Param("id"), user.(*models.User).ID, false, reaction.IsDisliked, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reaction updated successfully"})
}

// GetUserPosts retrieves posts created by a specific user.
func GetUserPosts(c *gin.Context) {
	userID := c.Param("userID")
	// Retrieve user-specific posts with a default limit for pagination
	posts, err := models.FindPostsByUserID(userID, c.Request.Context(), 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user's posts"})
		return
	}

	token, err := c.Cookie("token")
	if err != nil {
		token = ""
	}

	var authenticatedUserID primitive.ObjectID
	// Decode token to fetch user ID if token is present
	if token != "" {
		user, err := middleware.GetUserFromToken(token, c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode token"})
			log.Println(err)
			return
		}
		authenticatedUserID = user.ID
	}

	// Prepare posts for JSON response, enriching them with reaction data
	var postsResponse []map[string]interface{}
	for _, post := range posts {
		var likes, dislikes int
		var liked, disliked bool

		// Count likes and dislikes for each post and check user-specific reactions
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

		// Append structured post data to response slice
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
