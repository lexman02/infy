package routes

import (
	"infy/controllers"
	"infy/middleware"

	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the authentication routes for the application.
func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)                            // Handles user login
		auth.POST("/signup", controllers.Signup)                          // Handles user registration
		auth.GET("/user", middleware.Authorized(), controllers.User)      // Retrieves the logged-in user's profile
		auth.POST("/logout", middleware.Authorized(), controllers.Logout) // Handles user logout
	}
}
