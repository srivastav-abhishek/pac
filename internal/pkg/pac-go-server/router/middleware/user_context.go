package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
)

func fetchAndSetUserInfo(ctx *context.Context, accessToken string) error {
	user, err := kcClient.GetUserInfo(*ctx, accessToken, realm)
	if err != nil {
		return err
	}
	utils.SetContext(ctx, "username", *user.PreferredUsername)
	utils.SetContext(ctx, "userid", *user.Sub)

	return nil
}

func fetchAndSetUserGroups(ctx *context.Context, accessToken string) error {
	groups, err := kcClient.GetAccountGroups(*ctx, accessToken, realm)
	if err != nil {
		return fmt.Errorf("failed to get groups of a user: %w", err)
	}

	var grps []models.Group
	for _, group := range groups {
		grps = append(grps, models.Group{ID: *group.ID, Name: *group.Name})
	}

	utils.SetContext(ctx, "groups", grps)

	return nil
}

func fetchAndSetRoles(ctx *context.Context, accessToken string) error {
	_, claim, err := kcClient.DecodeAccessToken(*ctx, accessToken, realm)
	if err != nil {
		return fmt.Errorf("failed to parse the access token: %w", err)
	}

	getRoles := func() ([]string, error) {
		realmAccess, ok := (*claim)["realm_access"].(map[string]interface{})
		if !ok {
			return nil, errors.New("no realm_access in token")
		}
		roles, ok := realmAccess["roles"].([]interface{})
		if !ok {
			return nil, errors.New("no roles in realm_access")
		}
		var rolesList []string
		for _, role := range roles {
			rolesList = append(rolesList, role.(string))
		}
		return rolesList, nil
	}

	if roles, err := getRoles(); err != nil {
		return fmt.Errorf("failed to get roles of a user: %w", err)
	} else {
		utils.SetContext(ctx, "roles", roles)
	}

	return nil
}

func SetUserContexts(c *gin.Context) {
	ctx := c.Request.Context()

	utils.SetContext(&ctx, "keyclock_client", kcClient)
	utils.SetContext(&ctx, "keycloak_realm", realm)
	utils.SetContext(&ctx, "keycloak_hostname", hostname)

	accessToken := ctx.Value("keycloak_access_token").(string)

	if err := fetchAndSetUserInfo(&ctx, accessToken); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := fetchAndSetUserGroups(&ctx, accessToken); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := fetchAndSetRoles(&ctx, accessToken); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Replace the request context with the new context
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
