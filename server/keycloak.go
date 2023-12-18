package main

import (
	"fmt"
	"os"

	"github.com/Nerzal/gocloak/v13"
)

type keycloak struct {
	gocloak      gocloak.GoCloak // keycloak client
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
	realm        string          // realm specified in Keycloak
}

func newKeycloak() (*keycloak, error) {
	id := os.Getenv("KEYCLOAK_CLIENT_ID")
	secret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm := os.Getenv("KEYCLOAK_REALM")

	if id == "" || secret == "" || realm == "" {
		return nil, fmt.Errorf("missing keycloak environment variables: KEYCLOAK_CLIENT_ID or KEYCLOAK_CLIENT_SECRET or KEYCLOAK_REALM")
	}

	return &keycloak{
		gocloak:      *gocloak.NewClient("http://localhost:8080"),
		clientId:     id,
		clientSecret: secret,
		realm:        realm,
	}, nil
}
