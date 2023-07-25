package models

import "context"

// ExcludeGroups is a list of groups to exclude from the list of groups
var ExcludeGroups []string

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Membership is a flag to indicate if the user is a member of the group
	Membership bool     `json:"membership"`
	Quota      Capacity `json:"quota"`
}

func IsMemberOfGroup(ctx context.Context, name string) bool {
	groups := ctx.Value("groups").([]Group)
	for _, group := range groups {
		if group.Name == name {
			return true
		}
	}
	return false
}
