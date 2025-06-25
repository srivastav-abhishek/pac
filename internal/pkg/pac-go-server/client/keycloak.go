package client

import (
	"context"
	"errors"

	"github.com/Nerzal/gocloak/v13"
)

var (
	ErrorGroupNotFound = errors.New("group not found")
	ErrorUserNotFound  = errors.New("user not found")
)

// KeyCloackConfig holds the configuration for Keycloack operations
type KeyCloakConfig struct {
	Hostname    string
	AccessToken string
	Realm       string
	UserID      string
	Roles       []string
}

// KeycloackClient implements KeyCloackInterface
type KeyCloakClient struct {
	config KeyCloakConfig
	client *gocloak.GoCloak
}

func NewKeyCloakClient(config KeyCloakConfig) *KeyCloakClient {
	return &KeyCloakClient{
		config: config,
		client: gocloak.NewClient(config.Hostname),
	}
}

func (k *KeyCloakClient) GetClient() *gocloak.GoCloak {
	return k.client
}

// GetUsers for listing all the users from keycloak
func (k *KeyCloakClient) GetUsers(ctx context.Context) ([]*gocloak.User, error) {
	return k.client.GetUsers(ctx, k.config.AccessToken, k.config.Realm, gocloak.GetUsersParams{})
}

// GetUsers for listing all the users from keycloak
func (k *KeyCloakClient) GetUser(ctx context.Context, id string) (*gocloak.User, error) {
	return k.client.GetUser(ctx, k.config.AccessToken, k.config.Realm, id)
}

func (k *KeyCloakClient) GetGroups(ctx context.Context) ([]*gocloak.Group, error) {
	return k.client.GetGroups(ctx, k.config.AccessToken, k.config.Realm, gocloak.GetGroupsParams{})
}

func (k *KeyCloakClient) GetUserInfo(ctx context.Context) (*gocloak.UserInfo, error) {
	return k.client.GetUserInfo(ctx, k.config.AccessToken, k.config.Realm)
}

func (k *KeyCloakClient) AddUserToGroup(ctx context.Context, userID, groupID string) error {
	return k.client.AddUserToGroup(ctx, k.config.AccessToken, k.config.Realm, userID, groupID)
}

func (k *KeyCloakClient) DeleteUserFromGroup(ctx context.Context, userID, groupID string) error {
	return k.client.DeleteUserFromGroup(ctx, k.config.AccessToken, k.config.Realm, userID, groupID)
}

func (k *KeyCloakClient) GetUserGroups(ctx context.Context, userID string) ([]*gocloak.Group, error) {
	return k.client.GetUserGroups(ctx, k.config.AccessToken, k.config.Realm, userID, gocloak.GetGroupsParams{})
}

func (k *KeyCloakClient) DeleteUser(ctx context.Context, userID string) error {
	return k.client.DeleteUser(ctx, k.config.AccessToken, k.config.Realm, userID)
}

func (k *KeyCloakClient) IsRole(name string) bool {

	for _, role := range k.config.Roles {
		if role == name {
			return true
		}
	}
	return false
}

func (k *KeyCloakClient) GetUserID() string {
	return k.config.UserID
}

// GetConfigFromContext creates config from context
func GetConfigFromContext(ctx context.Context) KeyCloakConfig {
	config := KeyCloakConfig{
		Hostname:    ctx.Value("keycloak_hostname").(string),
		AccessToken: ctx.Value("keycloak_access_token").(string),
		Realm:       ctx.Value("keycloak_realm").(string),
	}

	if userID := ctx.Value("userid"); userID != nil {
		config.UserID = userID.(string)
	}

	if roles := ctx.Value("roles"); roles != nil {
		config.Roles = roles.([]string)
	}

	return config
}

var NewKeyCloakClientFromContext = func(ctx context.Context) Keycloak {
	config := GetConfigFromContext(ctx)
	return NewKeyCloakClient(config)
}
