package services

import (
	"context"

	"github.com/Nerzal/gocloak/v13"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/client"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
)

func getGroups(context context.Context) ([]*gocloak.Group, error) {
	groups, err := client.NewKeyClockClient(context).GetGroups()
	if err != nil {
		return nil, err
	}
	return filterExcludedGroups(groups), nil
}

func getGroup(context context.Context, id string) (*gocloak.Group, error) {
	groups, err := getGroups(context)
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if *group.ID == id {
			return group, nil
		}
	}
	return nil, client.ErrorGroupNotFound
}

// isExcludedGroup checks if the group is in the list of excluded groups
func isExcludedGroup(group string) bool {
	for _, excludedGroup := range models.ExcludeGroups {
		if group == excludedGroup {
			return true
		}
	}
	return false
}

// filterExcludedGroups filters out groups that are excluded
func filterExcludedGroups(groups []*gocloak.Group) []*gocloak.Group {
	var filteredGroups []*gocloak.Group
	for _, group := range groups {
		if !isExcludedGroup(*group.Name) {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}
