package controller

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Controller) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	accessToken := r.Header.Get("Authorization")

	user, err := c.Keycloak.Gocloak.GetUserInfo(ctx, accessToken, c.Keycloak.Realm)
	if err != nil {
		http.Error(w, prepareErrorResponse(err, "Error GetUserInfo"), http.StatusInternalServerError)
		return
	}

	me := Me{
		Username:      *user.PreferredUsername,
		FirstName:     *user.GivenName,
		LastName:      *user.FamilyName,
		FullName:      *user.Name,
		EmailVerified: *user.EmailVerified,
	}

	data, err := json.Marshal(me)
	if err != nil {
		http.Error(w, prepareErrorResponse(err, "Error Marshalling"), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}
