package routes

import (
	"github.com/gin-gonic/gin"
	"infy/controllers"
)

// AuthRoutes authentication routes
func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/signup", controllers.Signup)
	}
}
