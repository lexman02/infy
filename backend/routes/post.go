package routes

import (
	"github.com/gin-gonic/gin"
	"infy/controllers"
	"infy/middleware"
)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/posts")
	{
		post.GET("/", controllers.GetPosts)
		post.GET("/:id", controllers.GetPost)
		post.POST("/", middleware.Authorized(), controllers.CreatePost)
		post.PUT("/:id", middleware.Authorized(), controllers.UpdatePost)
		post.DELETE("/:id", middleware.Authorized(), controllers.DeletePost)
	}
}
