package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")
	admin.Use(middleware.Authorized())
	admin.Use(middleware.AdminAuthorized())
	{
		admin.GET("/users", controllers.GetUsers)
		admin.PUT("/users/:id", controllers.ToggleAdminStatus)
		admin.GET("/reports/posts", controllers.GetReportedPosts)
		admin.DELETE("/reports/posts/:id", controllers.DeleteReportedPost)
	}
}
