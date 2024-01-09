package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nerzal/gocloak/v13"
	"github.com/stretchr/testify/assert"

	"github.com/ar2rworld/projects-backend/keycloak"
)

func TestGetMe(t *testing.T) {
	authToken         := "token"
	testRealm         := "testRealm"
	testUsername      := "testUsername"
	testFirstName     := "testFirstName"
	testLastName      := "testLastName"
	testFullName      := "testFullName"
	testEmailVerified := true

	k := &keycloak.Keycloak{ Realm: testRealm }

	mockCloak := &keycloak.GoCloaklikeMock{
		GetUserInfoFunc: func(ctx context.Context, accessToken, realm string) (*gocloak.UserInfo, error) {
			assert.Equal(t, realm, testRealm, "Incorrect realm")

			return &gocloak.UserInfo{
				PreferredUsername: &testUsername,
				GivenName: &testFirstName,
				FamilyName: &testLastName,
				Name: &testFullName,
				EmailVerified: &testEmailVerified,
			}, nil
		},
	}
	
	k.Gocloak = mockCloak

	c := NewController(k)

	r, _  := http.NewRequest(http.MethodGet, "/me", nil)
	r.Header.Add("Authorization", authToken)

	p := httptest.NewRecorder()
	
	c.GetMe(p, r)

	assert.Equal(t, p.Code, http.StatusOK, "Status code is not Ok")

	var data Me
	err := json.Unmarshal(p.Body.Bytes(), &data)
	
	assert.NoError(t, err, "Error unmarshalling response")

	assert.Equal(t, data.Username, testUsername, "Username is not correct");
	
	assert.Equal(t, data.FirstName, testFirstName, "FirstName is not correct");

	assert.Equal(t, data.LastName, testLastName, "LastName is not correct");

	assert.Equal(t, data.FullName, testFullName, "FullName is not correct");

	assert.Equal(t, data.EmailVerified, testEmailVerified, "EmailVerified is not correct");
}
