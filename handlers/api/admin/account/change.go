package account

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/models/admin/account"
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
)

// ChangePassword controller for change password of account
func ChangePassword(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("values").(middlewares.Vars)
		_, err := middlewares.FetchAuth(userRole.AccessUuid, client)
		if err != nil {
			response.RespondWithError(w, http.StatusUnauthorized, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
			return
		}
		req := account.ChangeAccountRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.ChangeAccountValidate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		changePassword, err := service.ChangePassword(req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, changePassword)
	}
}
