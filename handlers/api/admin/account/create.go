package account

import (
	"api-trainning-center/models/admin/account"
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

const ADMIN = "ADMIN"

// CreateAccount controller for creating new users
func CreateAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := account.AccountRequest{}
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

		resp, err := service.CreateAccount(req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
