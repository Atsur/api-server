package middlewares

import (
	"net/http"
	"log"
	"time"
	"strings"
	"firebase.google.com/go/auth"
	"github.com/atsur/api-server/pkg/crypto"
	"github.com/gin-gonic/gin"
	"github.com/atsur/api-server/pkg/secrets"
)

const (
	authorizationHeader = "Authorization"
	apiKeyHeader        = "X-API-Key"
	cronExecutedHeader  = "X-Appengine-Cron"
	valName             = "FIREBASE_ID_TOKEN"
)


func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}

// Gin middleware for JWT auth
func AuthJWT(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		authHeader := c.Request.Header.Get(authorizationHeader)
		log.Println("authHeader", authHeader)
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		log.Println("token = ", token)
		idToken, err := client.VerifyIDToken(c, token) // usually hits a local cache
		if err != nil {
			log.Println("err = ", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("Auth time:", time.Since(startTime))

		c.Set(valName, idToken)
		c.Next()
	}
}

// API key auth middleware
func AuthAPIKey(secretId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get(apiKeyHeader)
		log.Println("header key is = ",key)
		secret, err := secrets.GetSecret(secretId)
		if err != nil {
			log.Println("failed to get secret = ",err)
			log.Println("failed to get secret err = ",err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("comparing secret with provided key", secret, key)

		if secret != key {
			log.Println("key doesnt match!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		log.Println("no error during check")
		c.Next()
	}
}
