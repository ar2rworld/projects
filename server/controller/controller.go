package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	
	"github.com/ar2rworld/projects-backend/keycloak"
)

type doc struct {
	Id   string    `json:"id"`
	Num  string    `json:"num"`
	Date time.Time `json:"date"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Ok bool `json:"ok"`
	Description string `json:"description"`
	Message string `json:"message"`
	
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`

	Username string
	FirstName string
	LastName string
	FullName string
	EmailVerified bool
}

type Me struct {
	Username string
	FirstName string
	LastName string
	FullName string
	EmailVerified bool
}

type Controller struct {
	Keycloak *keycloak.Keycloak
}

func NewController(keycloak *keycloak.Keycloak) *Controller {
	return &Controller{
		Keycloak: keycloak,
	}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rq := &LoginRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, prepareErrorResponse(err, "username and password required"), http.StatusBadRequest)
		return
	}

	jwt, err := c.Keycloak.Gocloak.Login(ctx,
		c.Keycloak.ClientId,
		c.Keycloak.ClientSecret,
		c.Keycloak.Realm,
		rq.Username,
		rq.Password)

	if err != nil {
		http.Error(w, prepareErrorResponse(err, "Invalid credentials"), http.StatusForbidden)
		return
	}

	_, m, err := c.Keycloak.Gocloak.DecodeAccessToken(ctx, jwt.AccessToken, c.Keycloak.Realm)
	if err != nil {
		log.Println(err)
		http.Error(w, prepareErrorResponse(err, "Unable to decode your token"), http.StatusForbidden)
		return
	}

	rs := &LoginResponse{
		Ok:           true,
		Description:  "Success",
		Message:      "Success",
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,

		Username: (*m)["preferred_username"].(string),
		FirstName: (*m)["given_name"].(string),
		LastName: (*m)["family_name"].(string),
		FullName: (*m)["name"].(string),
		EmailVerified: (*m)["email_verified"].(bool),
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}

func (c *Controller) getDocs(w http.ResponseWriter, r *http.Request) {

	rs := []*doc{
		{
			Id:   "1",
			Num:  "ABC-123",
			Date: time.Now().UTC(),
		},
		{
			Id:   "2",
			Num:  "ABC-456",
			Date: time.Now().UTC(),
		},
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)

}
