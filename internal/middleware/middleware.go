package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/arashiaslan/forum-app-go/internal/configs"
	"github.com/arashiaslan/forum-app-go/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		rawHeader := c.Request.Header.Get("Authorization")
		// log raw header for debugging (prints quoted string)
		log.Printf("AUTH RAW HEADER: %q", rawHeader)

		// Normalize: trim spaces, remove Bearer prefix, remove surrounding quotes
		h := strings.TrimSpace(rawHeader)
		if h == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Remove case-insensitive "Bearer " prefix if present
		if strings.HasPrefix(strings.ToLower(h), "bearer ") {
			h = strings.TrimSpace(h[len("Bearer "):])
		}

		// Remove wrapping quotes if any
		h = strings.Trim(h, `"'`)

		// log normalized token and bytes for debugging
		log.Printf("AUTH NORMALIZED TOKEN (len=%d): %q", len(h), h)
		log.Printf("AUTH TOKEN BYTES: % x", []byte(h))

		if h == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		userID, username, err := jwt.ValidateToken(h, secretKey)
		if err != nil {
			// log error to help find cause of malformed token
			log.Printf("AUTH: token validate error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}