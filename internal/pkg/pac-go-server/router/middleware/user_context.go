package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/utils"
)

func fetchAndSetUserInfo(c *gin.Context, ctx context.Context, accessToken string) {
	user, err := kcClient.GetUserInfo(ctx, accessToken, realm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.SetContext(&ctx, "username", *user.PreferredUsername)
	utils.SetContext(&ctx, "userid", *user.Sub)
}

func fetchAndSetUserGroups(c *gin.Context, ctx context.Context, accessToken string) {
	groups, err := kcClient.GetAccountGroups(ctx, accessToken, realm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to get groups of a user: %s", err.Error())})
		return
	}

	var grps []models.Group
	for _, group := range groups {
		grps = append(grps, models.Group{ID: *group.ID, Name: *group.Name})
	}

	utils.SetContext(&ctx, "groups", grps)
}

func fetchAndSetRoles(c *gin.Context, ctx context.Context, accessToken string) {
	_, claim, err := kcClient.DecodeAccessToken(ctx, accessToken, realm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to parse the access token: %s", err.Error())})
		return
	}

	getRoles := func() ([]string, error) {
		realmAccess, ok := (*claim)["realm_access"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("no realm_access in token")
		}
		roles, ok := realmAccess["roles"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("no roles in realm_access")
		}
		var rolesList []string
		for _, role := range roles {
			rolesList = append(rolesList, role.(string))
		}
		return rolesList, nil
	}

	if roles, err := getRoles(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to get roles of a user: %s", err.Error())})
		return
	} else {
		utils.SetContext(&ctx, "roles", roles)
	}
}

func SetUserContexts(c *gin.Context) {
	ctx := c.Request.Context()

	utils.SetContext(&ctx, "keyclock_client", kcClient)
	utils.SetContext(&ctx, "keycloak_realm", realm)
	utils.SetContext(&ctx, "keycloak_hostname", hostname)

	accessToken := ctx.Value("keycloak_access_token").(string)

	fetchAndSetUserInfo(c, ctx, accessToken)

	fetchAndSetUserGroups(c, ctx, accessToken)

	fetchAndSetRoles(c, ctx, accessToken)

	// Replace the request context with the new context
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
