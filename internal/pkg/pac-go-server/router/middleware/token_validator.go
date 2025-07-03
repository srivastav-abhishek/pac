package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
)

func fetchBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) < 1 {
		return "", errors.New("authorization header is missing")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("authorization header format must be 'Bearer <token>'")
	}

	token := strings.TrimSpace(authHeader[len(prefix):])
	if token == "" {
		return "", errors.New("token is empty")
	}

	return token, nil
}

func ValidateToken(c *gin.Context) {
	ctx := c.Request.Context()

	accessToken, err := fetchBearerToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	utils.SetContext(&ctx, "keycloak_access_token", accessToken)
	rptResult, err := kcClient.RetrospectToken(ctx, accessToken, clientId, clientSecret, realm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	isTokenValid := *rptResult.Active
	if !isTokenValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid. Please fetch a new token and try again!"})
		return
	}

	// Replace the request context with the new context
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
