package services

import (
	"fmt"
	"net/http"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/client"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	config := client.GetConfigFromContext(c)
	usrs, err := client.NewKeyCloakClient(config).GetUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var users []models.User
	for _, user := range usrs {
		u := models.User{
			Username: *user.Username,
			ID:       *user.ID,
		}
		if user.FirstName != nil {
			u.FirstName = *user.FirstName
		}
		if user.LastName != nil {
			u.LastName = *user.LastName
		}
		// Email field is optional, hence check for nil before assigning.
		if user.Email != nil {
			u.Email = *user.Email
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	config := client.GetConfigFromContext(c)
	usrs, err := client.NewKeyCloakClient(config).GetUsers(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, user := range usrs {
		if *user.ID == id {
			usr := models.User{
				Username:  *user.Username,
				ID:        *user.ID,
				FirstName: *user.FirstName,
				LastName:  *user.LastName,
			}
			if user.FirstName != nil {
				usr.FirstName = *user.FirstName
			}
			if user.LastName != nil {
				usr.LastName = *user.LastName
			}
			// Email field is optional, hence check for nil before assigning.
			if user.Email != nil {
				usr.Email = *user.Email
			}
			c.JSON(http.StatusOK, usr)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("user: %s not found", id)})
}
