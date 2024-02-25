package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Initialize the routes
	AuthRoutes(router)
	PostRoutes(router)

	return router
}
