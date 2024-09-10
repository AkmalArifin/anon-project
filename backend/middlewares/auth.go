package middlewares

import (
	"net/http"
	"strings"

	"example.com/anon-project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	// Split the header value into parts
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	token := parts[1]

	id, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorize"})
		return
	}

	c.Set("user_id", id)
	c.Next()
}
