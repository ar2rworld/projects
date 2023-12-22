package keycloak

import (
	"fmt"
	"os"

	"github.com/Nerzal/gocloak/v13"
)

type Keycloak struct {
	Gocloak      gocloak.GoCloak // keycloak client
	ClientId     string          // clientId specified in Keycloak
	ClientSecret string          // client secret specified in Keycloak
	Realm        string          // realm specified in Keycloak
}

func NewKeycloak() (*Keycloak, error) {
	id := os.Getenv("KEYCLOAK_CLIENT_ID")
	secret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm := os.Getenv("KEYCLOAK_REALM")

	if id == "" || secret == "" || realm == "" {
		return nil, fmt.Errorf("missing keycloak environment variables: KEYCLOAK_CLIENT_ID or KEYCLOAK_CLIENT_SECRET or KEYCLOAK_REALM")
	}

	return &Keycloak{
		Gocloak:      *gocloak.NewClient("http://localhost:8080"),
		ClientId:     id,
		ClientSecret: secret,
		Realm:        realm,
	}, nil
}
