package middleware

import (
	"context"
	"infy/models"
	"infy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Authorized checks if the user is authorized by verifying the JWT token.
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token") // Attempt to retrieve the JWT token from the cookie.
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"}) // Send an unauthorized error if no token is found.
			c.Abort()
			return
		}

		user, err := GetUserFromToken(token, c.Request.Context()) // Decode the token to retrieve the user.
		if err != nil || user == nil {
			c.JSON(401, gin.H{"error": "Unauthorized"}) // Respond with unauthorized if the token is invalid or user does not exist.
			c.Abort()
			return
		}

		c.Set("user", user) // Set the user in the context for downstream handlers.
		c.Next()
	}
}

// GetUserFromToken decodes a JWT token and retrieves the user from the database.
func GetUserFromToken(token string, ctx context.Context) (*models.User, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		jwtSecretKey := utils.GetEnv("JWT_SECRET_KEY", "") // Retrieve JWT secret key from environment.
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err // Return error if the token parsing fails.
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userId := claims["sub"].(string)        // Extract user ID from token claims.
		return models.FindUserByID(userId, ctx) // Fetch the user from the database by ID.
	}

	return nil, nil // Return nil if the token is not valid or claims are not correctly parsed.
}

// ErrorHandler captures any errors occurred during HTTP request processing and returns them.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request.
		if len(c.Errors) > 0 {
			c.JSON(400, gin.H{"errors": c.Errors}) // Return any errors caught during request processing.
			c.Abort()
		}
	}
}

// AdminAuthorized checks if the logged-in user has admin privileges.
func AdminAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists || !user.(*models.User).IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin privileges required"}) // Check if the user is an admin.
			c.Abort()
			return
		}
		c.Next()
	}
}
