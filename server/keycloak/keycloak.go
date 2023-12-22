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
	kcHost := os.Getenv("KEYCLOAK_HOST")

	if id == "" || secret == "" || realm == "" || kcHost == ""  {
		return nil, fmt.Errorf(`missing keycloak environment variables:
		KEYCLOAK_CLIENT_ID=%s or
		KEYCLOAK_CLIENT_SECRET=%s or
		KEYCLOAK_REALM=%s or
		KEYCLOAK_HOST=%s
		`, id, secret, realm, kcHost)
	}

	return &Keycloak{
		Gocloak:      *gocloak.NewClient(kcHost),
		ClientId:     id,
		ClientSecret: secret,
		Realm:        realm,
	}, nil
}
