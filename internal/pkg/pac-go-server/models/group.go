package models

// ExcludeGroups is a list of groups to exclude from the list of groups
var ExcludeGroups []string

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Membership is a flag to indicate if the user is a member of the group
	Membership bool     `json:"membership"`
	Quota      Capacity `json:"quota"`
}

// isExcludedGroup checks if the group is in the list of excluded groups
func isExcludedGroup(group string) bool {
	for _, excludedGroup := range ExcludeGroups {
		if group == excludedGroup {
			return true
		}
	}
	return false
}

// FilterExcludedGroups filters out groups that are excluded
func FilterExcludedGroups(groups []Group) []Group {
	var filteredGroups []Group
	for _, group := range groups {
		if !isExcludedGroup(group.Name) {
			filteredGroups = append(filteredGroups, group)
		}
	}
	return filteredGroups
}
