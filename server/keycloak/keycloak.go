package keycloak

import (
	"context"
	"fmt"
	"os"

	gocloak "github.com/Nerzal/gocloak/v13"
	jwt "github.com/golang-jwt/jwt/v4"
)

//go:generate moq -out gocloaklike_moq.go . GoCloaklike

type GoCloaklike interface {
	CreateUser(ctx context.Context, token, realm string, user gocloak.User) (string, error)
	DecodeAccessToken(ctx context.Context, accessToken string, realm string) (*jwt.Token, *jwt.MapClaims, error)
	GetUserInfo(ctx context.Context, accessToken, realm string) (*gocloak.UserInfo, error)
	GetUsers(ctx context.Context, token, realm string, params gocloak.GetUsersParams) ([]*gocloak.User, error)
	Login(ctx context.Context, clientId string, clientSecret string, realm string, username string, password string) (*gocloak.JWT, error)
	LoginAdmin(ctx context.Context, username, password, realm string) (*gocloak.JWT, error)
	SetPassword(ctx context.Context, token, userID, realm, password string, temporary bool) error
}

type Keycloak struct {
	Gocloak       GoCloaklike 		// keycloak client
	ClientId   	  string          // clientId specified in Keycloak
	ClientSecret  string          // client secret specified in Keycloak
	Realm     	  string          // realm specified in Keycloak
	AdminUsername string			    // admin username
	AdminPassword string			    // admin password
}

func NewKeycloak() (*Keycloak, error) {
	id := os.Getenv("KEYCLOAK_CLIENT_ID")
	secret := os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm := os.Getenv("KEYCLOAK_REALM")
	kcHost := os.Getenv("KEYCLOAK_HOST")
	kcAdminUsername := os.Getenv("KEYCLOAK_ADMIN_USERNAME")
	kcAdminPassword := os.Getenv("KEYCLOAK_ADMIN_PASSWORD")

	if id == "" || secret == "" || realm == "" || kcHost == "" ||
	kcAdminUsername == "" || kcAdminPassword == "" {
		return nil, fmt.Errorf(`missing keycloak environment variables:
		KEYCLOAK_CLIENT_ID=%s or
		KEYCLOAK_CLIENT_SECRET=%s or
		KEYCLOAK_REALM=%s or
		KEYCLOAK_HOST=%s or
		KEYCLOAK_ADMIN_USERNAME=%s or
		KEYCLOAK_ADMIN_PASSWORD=%s
		`, id, secret, realm, kcHost, kcAdminUsername, kcAdminPassword)
	}

	return &Keycloak{
		Gocloak:       gocloak.NewClient(kcHost),
		ClientId:      id,
		ClientSecret:  secret,
		Realm:         realm,
		AdminUsername: kcAdminUsername,
		AdminPassword: kcAdminPassword,
	}, nil
}
