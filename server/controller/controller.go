package controller

import (
	"encoding/json"
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

type RegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	FavoriteAnimal string `json:"favanimal"`
}
type RegisetResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
	Message     string `json:"message"`
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
