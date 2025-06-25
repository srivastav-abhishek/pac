package client

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

//go:generate mockgen -destination=mock_keycloak.go -package=client . Keycloak
type Keycloak interface {
	GetClient() *gocloak.GoCloak
	GetUsers(ctx context.Context) ([]*gocloak.User, error)
	GetUser(ctx context.Context, id string) (*gocloak.User, error)
	GetGroups(ctx context.Context) ([]*gocloak.Group, error)
	GetUserInfo(ctx context.Context) (*gocloak.UserInfo, error)
	AddUserToGroup(ctx context.Context, userID, groupID string) error
	DeleteUserFromGroup(ctx context.Context, userID, groupID string) error
	GetUserGroups(ctx context.Context, userID string) ([]*gocloak.Group, error)
	DeleteUser(ctx context.Context, userID string) error
	IsRole(name string) bool
	GetUserID() string
}
