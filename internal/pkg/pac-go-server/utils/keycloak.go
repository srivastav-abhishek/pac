package utils

import (
	"context"
	"errors"

	"github.com/Nerzal/gocloak/v13"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
)

var (
	ErrorGroupNotFound = errors.New("group not found")
	ErrorUserNotFound  = errors.New("user not found")
)

const (
	ManagerRole = "manager"
)

type KeyClockClient struct {
	ctx         context.Context
	client      *gocloak.GoCloak
	accessToken string
	realm       string
	hostname    string
}

func NewKeyClockClient(ctx context.Context) *KeyClockClient {
	hostname := ctx.Value("keycloak_hostname").(string)

	return &KeyClockClient{
		ctx:         ctx,
		client:      gocloak.NewClient(hostname),
		accessToken: ctx.Value("keycloak_access_token").(string),
		realm:       ctx.Value("keycloak_realm").(string),
	}
}

func (k *KeyClockClient) GetClient() *gocloak.GoCloak {
	return k.client
}

// GetUser by ID
func (k *KeyClockClient) GetUser(id string) (*models.User, error) {
	users, err := k.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, ErrorUserNotFound
}

// GetUsers for listing all the users from keycloak
func (k *KeyClockClient) GetUsers() ([]models.User, error) {
	var users []models.User
	usr, err := k.client.GetUsers(k.ctx, k.accessToken, k.realm, gocloak.GetUsersParams{})
	if err != nil {
		return nil, err
	}
	for _, user := range usr {
		u := models.User{
			Username:  *user.Username,
			ID:        *user.ID,
			FirstName: *user.FirstName,
			LastName:  *user.LastName,
		}
		// Email field is optional, hence check for nil before assigning.
		if user.Email != nil {
			u.Email = *user.Email
		}
		users = append(users, u)
	}
	return users, nil
}

func (k *KeyClockClient) GetGroups() ([]models.Group, error) {
	var groups []models.Group
	grp, err := k.client.GetGroups(k.ctx, k.accessToken, k.realm, gocloak.GetGroupsParams{})
	if err != nil {
		return nil, err
	}
	for _, group := range grp {
		groups = append(groups, models.Group{
			Name: *group.Name,
			ID:   *group.ID,
		})
	}
	return groups, nil
}

func (k *KeyClockClient) GetGroup(id string) (*models.Group, error) {
	groups, err := k.GetGroups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group.ID == id {
			return &group, nil
		}
	}
	return nil, ErrorGroupNotFound
}

func (k *KeyClockClient) AddUserToGroup(userID, groupID string) error {
	return k.client.AddUserToGroup(k.ctx, k.accessToken, k.realm, userID, groupID)
}

func (k *KeyClockClient) IsMemberOfGroup(name string) bool {
	groups := k.GetUserGroups()
	for _, group := range groups {
		if group.Name == name {
			return true
		}
	}
	return false
}

func (k *KeyClockClient) IsRole(name string) bool {
	roles := k.ctx.Value("roles").([]string)

	for _, role := range roles {
		if role == name {
			return true
		}
	}
	return false
}

func (k *KeyClockClient) GetUserID() string {
	return k.ctx.Value("userid").(string)
}

func (k *KeyClockClient) GetUserGroups() []models.Group {
	return k.ctx.Value("groups").([]models.Group)
}
