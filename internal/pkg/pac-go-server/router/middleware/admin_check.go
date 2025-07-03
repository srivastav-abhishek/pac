package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pacClient "github.com/PDeXchange/pac/internal/pkg/pac-go-server/client"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
)

func AllowAdminOnly(c *gin.Context) {
	kc := pacClient.NewKeyClockClient(c.Request.Context())
	if !kc.IsRole(utils.ManagerRole) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized to perform this action"})
		return
	}
	c.Next()
}
