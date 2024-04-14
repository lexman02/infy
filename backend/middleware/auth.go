package middleware

import (
	"infy/models"
	"infy/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Authorized checks if the user is authorized by checking the JWT token
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parse the token with the JWT_SECRET_KEY from the environment variables
		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			jwtSecretKey := utils.GetEnv("JWT_SECRET_KEY", "")
			return []byte(jwtSecretKey), nil
		})
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(401, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		// Get the user from the token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["sub"].(string)
			user, err := models.FindUserByID(userId, c.Request.Context())
			if err != nil {
				c.JSON(401, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
			c.Set("user", user)
			c.Next()
		}
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(400, gin.H{"errors": c.Errors})
			c.Abort()
		}
	}
}

func AdminAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists || !user.(*models.User).IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin privileges required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
