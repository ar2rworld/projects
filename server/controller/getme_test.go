package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ar2rworld/projects-backend/keycloak"
)

func TestGetMe(t *testing.T) {
	k := &keycloak.Keycloak{}
	
	// k := &MyKeycloak{}
	c := NewController(k)

	authToken         := "Bearer token"
	testUsername      := "testUsername"
	testFirstName     := "testUsername"
	testLastName      := "testUsername"
	testFullName      := "testUsername"
	testEmailVerified := true

	r, _  := http.NewRequest(http.MethodGet, "/me", nil)
	r.Header.Add("Authorization", authToken)

	p := httptest.NewRecorder()
	
	c.GetMe(p, r)

	if p.Code != http.StatusOK {
		t.Fatal("Status code is not Ok")
	}

	var data Me
	err := json.Unmarshal(p.Body.Bytes(), &data)
	if err != nil {
		t.Fatal("Error unmarshalling response:", string(p.Body.Bytes()))
	}

	if data.Username != testUsername {
		t.Fatal("Username is not correct")
	}
	if data.FirstName != testFirstName {
		t.Fatal("FirstName is not correct")
	}
	if data.LastName != testLastName {
		t.Fatal("LastName is not correct")
	}
	if data.FullName != testFullName {
		t.Fatal("FullName is not correct")
	}
	if data.EmailVerified != testEmailVerified {
		t.Fatal("EmailVerified is not correct")
	}
}

type MyKeycloak struct {

	keycloak.Keycloak
}
