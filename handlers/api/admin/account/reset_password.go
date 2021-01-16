package account

import (
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
)

type ResetPasswordRequest struct {
	UserName string `json:"username"`
}

// ResetPassword controller for account just have Admin
func ResetPassword(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := ResetPasswordRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if req.UserName == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Bạn chưa nhập tên đăng nhập"))
			return
		}

		resetPassword, err := service.ResetPassword(req.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resetPassword)
	}
}
