package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// AdminRoutes defines routes that are only accessible by users with administrative privileges.
func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.Authorized())      // Requires authorization token
	admin.Use(middleware.AdminAuthorized()) // Requires admin-level access
	{
		admin.GET("/users", controllers.GetUsers)                          // Retrieves all users
		admin.PUT("/users/:id", controllers.ToggleAdminStatus)             // Toggles admin status of a user
		admin.GET("/reports/posts", controllers.GetReportedPosts)          // Retrieves reported posts
		admin.DELETE("/reports/posts/:id", controllers.DeleteReportedPost) // Deletes a reported post
	}
}
