package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type doc struct {
	Id   string    `json:"id"`
	Num  string    `json:"num"`
	Date time.Time `json:"date"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
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

type controller struct {
	keycloak *keycloak
}

func newController(keycloak *keycloak) *controller {
	return &controller{
		keycloak: keycloak,
	}
}

func (c *controller) login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rq := &loginRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, prepareErrorResponse(err, "username and password required"), http.StatusBadRequest)
		return
	}

	jwt, err := c.keycloak.gocloak.Login(ctx,
		c.keycloak.clientId,
		c.keycloak.clientSecret,
		c.keycloak.realm,
		rq.Username,
		rq.Password)

	if err != nil {
		http.Error(w, prepareErrorResponse(err, "Invalid credentials"), http.StatusForbidden)
		return
	}

	_, m, err := c.keycloak.gocloak.DecodeAccessToken(ctx, jwt.AccessToken, c.keycloak.realm)
	if err != nil {
		log.Println(err)
		http.Error(w, prepareErrorResponse(err, "Unable to decode your token"), http.StatusForbidden)
		return
	}

	rs := &loginResponse{
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

func (c *controller) getDocs(w http.ResponseWriter, r *http.Request) {

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
