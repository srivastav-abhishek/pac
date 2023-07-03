package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	log "github.com/PDeXchange/pac/internal/pkg/pac-go-server/logger"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
)

func GetAllGroups(c *gin.Context) {
	logger := log.GetLogger()
	var err error
	groups := []models.Group{}
	groups, err = utils.NewKeyClockClient(c.Request.Context()).GetGroups()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, group := range groups {
		groups[i].Membership = utils.NewKeyClockClient(c.Request.Context()).IsMemberOfGroup(group.Name)
	}

	logger.Debug("fetching groups quota")
	var groupIds []string
	for _, grp := range groups {
		groupIds = append(groupIds, grp.ID)
	}
	groupsQuota, err := dbCon.GetGroupsQuota(groupIds)
	if err != nil {
		logger.Error("failed to get groups quota", zap.Any("group ids", groupIds), zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Debug("fetched groups quota", zap.Any("groups quota", groupsQuota))

	// construct a map with quota associated for group
	quotaMap := make(map[string]models.Capacity)
	for _, groupQuota := range groupsQuota {
		quotaMap[groupQuota.GroupID] = groupQuota.Capacity
	}

	// update groups quota
	for index, group := range groups {
		if quota, ok := quotaMap[group.ID]; ok {
			groups[index].Quota = quota
		}
	}
	logger.Debug("groups with quota", zap.Any("groups", groups))
	c.JSON(http.StatusOK, groups)
}

func GetGroup(c *gin.Context) {
	id := c.Param("id")
	grp, err := utils.NewKeyClockClient(c.Request.Context()).GetGroup(id)
	if err != nil && err != utils.ErrorGroupNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if err == utils.ErrorGroupNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grp)
}
