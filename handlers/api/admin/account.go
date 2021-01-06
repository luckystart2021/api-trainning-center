package admin

import (
	"api-trainning-center/models"
	"api-trainning-center/service/account"
	"api-trainning-center/service/response"
	"errors"

	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

const (
	ADMIN = "ADMIN"
	ROLE  = "Role"
)

// CreateAccount zcontroller for creating new users
func CreateAccount(service account.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value(ROLE)
		if userRole != ADMIN {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("You don't have permission to sign up"))
			return
		}
		req := models.AccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.Validate(""); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		if _, err := req.IsValid(); err != nil {
			// Check Role have is valid
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := service.CreateAccount(req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func Login(service account.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.AccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.Validate("login"); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		login, err := service.Login(req, client)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, login)
	}
}
