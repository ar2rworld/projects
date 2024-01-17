package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

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
