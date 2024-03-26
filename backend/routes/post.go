package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.GET("/", controllers.GetPosts)
		post.GET("/:id", controllers.GetPost)
		post.POST("/", middleware.Authorized(), controllers.CreatePost)
		post.PUT("/:id", middleware.Authorized(), controllers.UpdatePost)
		post.DELETE("/:id", middleware.Authorized(), controllers.DeletePost)
		post.GET("/user/:userID", controllers.GetUserPosts) // New route to fetch posts by a specific user
	}
}
