package services

import (
	"net/http"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := utils.NewKeyClockClient(c.Request.Context()).GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	usr, err := utils.NewKeyClockClient(c.Request.Context()).GetUser(id)
	if err != nil && err != utils.ErrorUserNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err == utils.ErrorUserNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usr)
}
