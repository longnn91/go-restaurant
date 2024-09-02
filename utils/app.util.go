package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserIdFromContext retrieves the user ID from the Gin context.
// It returns the user ID as a string and an error if the user ID is not found.
func GetUserIdFromContext(c *gin.Context) (int, error) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return 0, errors.New("user ID not found in context")
	}

	// Type assert userId to an int
	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		c.Abort()
		return 0, errors.New("user ID has invalid format")
	}

	return userIdInt, nil
}
