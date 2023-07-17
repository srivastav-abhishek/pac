package services

import (
	"net/http"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
	"github.com/gin-gonic/gin"
)

// GetEvents returns all events
func GetEvents(c *gin.Context) {
	kc := utils.NewKeyClockClient(c.Request.Context())

	var userID string
	if !kc.IsRole(utils.ManagerRole) {
		// Get authenticated user's ID
		userID = kc.GetUserID()
	}
	events, err := dbCon.GetEventsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
