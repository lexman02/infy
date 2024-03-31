package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// AuthRoutes authentication routes
func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/signup", controllers.Signup)
		auth.GET("/user", middleware.Authorized(), controllers.User)
		auth.POST("/logout", middleware.Authorized(), controllers.Logout)
	}
}
