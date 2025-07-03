package middleware

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

var (
	clientId     = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm        = os.Getenv("KEYCLOAK_REALM")
	hostname     = os.Getenv("KEYCLOAK_HOSTNAME")
	kcClient     *gocloak.GoCloak
)

func init() {
	kcClient = gocloak.NewClient(hostname)
}
