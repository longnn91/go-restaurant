package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (string, bool) {
	// Parse the token with the secret key
	secretKey := os.Getenv("SECRET_KEY")
	jwtKey := []byte(secretKey)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})

	// Check for parsing errors
	if err != nil {
		return "", false
	}

	// Check if the token is valid
	if !token.Valid {
		return "", false
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	// Extract the user ID from the claims
	var userId string
	switch v := claims["userId"].(type) {
	case string:
		userId = v
	case float64:
		userId = fmt.Sprintf("%.0f", v) // Convert float64 to string, assuming it's an integer
	default:
		return "", false
	}

	// Return the user ID and verification status
	return userId, true
}
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userId, byPass := verifyToken(authHeader)

	if !byPass {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("userId", userId)

	c.Next()
}
