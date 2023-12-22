package keycloak

import (
	"context"
	"fmt"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v4"
)

type GoCloaklike interface {
	Login(ctx context.Context, clientId string, clientSecret string, realm string, username string, password string) (*gocloak.JWT, error)
	DecodeAccessToken(ctx context.Context, accessToken string, realm string) (*jwt.Token, *jwt.MapClaims, error)
}

type Keycloak struct {
	Gocloak      GoCloaklike 		 // keycloak client
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
		Gocloak:      gocloak.NewClient(kcHost),
		ClientId:     id,
		ClientSecret: secret,
		Realm:        realm,
	}, nil
}
