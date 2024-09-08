package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserIdFromContext retrieves the user ID from the Gin context.
// It returns the user ID as a string and an error if the user ID is not found.
func GetUserIdFromContext(c *gin.Context) (int, error) {
	userId, _ := c.Get("userId")

	// Type assert userId to an int
	userIdInt, err := strconv.Atoi(userId.(string))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		c.Abort()
		return 0, errors.New("user ID has invalid format")
	}

	return userIdInt, nil
}
