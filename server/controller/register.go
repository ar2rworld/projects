package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
)

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rq := &RegisterRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(rq); err != nil {
		http.Error(w, prepareErrorResponse(err, "username and password required"), http.StatusBadRequest)
		return
	}

	jwt, err := c.Keycloak.Gocloak.LoginAdmin(ctx, c.Keycloak.AdminUsername, c.Keycloak.AdminPassword, c.Keycloak.Realm)
	if err != nil {
		log.Println(err)
		http.Error(w, prepareErrorResponse(err, "InternalServerError"), http.StatusInternalServerError)
		return
	}

	users, err := c.Keycloak.Gocloak.GetUsers(
		ctx,
		jwt.AccessToken,
		c.Keycloak.Realm,
		gocloak.GetUsersParams{Username: &rq.Username},
	)

	if err != nil {
		log.Println(err)
		http.Error(w, prepareErrorResponse(err, "error GetUsers"), http.StatusBadRequest)
		return
	}

	if len(users) >= 1 {
		log.Println("user already exists")
		http.Error(w, prepareErrorResponse(nil, "user already exists"), http.StatusBadRequest)
		return
	}

	enabled := true
	
	u := gocloak.User{
		Enabled:  &enabled,
		Username: &rq.Username,
	}

	// add favorite pet to attributes
	if rq.FavoriteAnimal != "" {
		u.Attributes = &map[string][]string{"favanimal": {rq.FavoriteAnimal}}
	}

	userID, err := c.Keycloak.Gocloak.CreateUser(
		ctx,
		jwt.AccessToken,
		c.Keycloak.Realm,
		u,
	)

	if err != nil {
		log.Println(err)
		http.Error(w, prepareErrorResponse(err, "Error creating a user"), http.StatusForbidden)
		return
	}

	err = c.Keycloak.Gocloak.SetPassword(ctx, jwt.AccessToken, userID, c.Keycloak.Realm, rq.Password, false)
	if err != nil {
		log.Println(err)
		http.Error(w, prepareErrorResponse(err, "could not set password"), http.StatusForbidden)
		return
	}

	rs := &RegisetResponse{
		Ok:           true,
		Description:  "Success registering user",
		Message:      "Success",
	}

	rsJs, _ := json.Marshal(rs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(rsJs)
}
