package admin

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/models/admin"

	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
)

const ADMIN = "ADMIN"

type ResetPasswordRequest struct {
	UserName string `json:"username"`
}

// CreateAccount controller for creating new users
func CreateAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := admin.AccountRequest{}
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

// LogoutAccount controller for logout
func LogoutAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Context().Value("values").(middlewares.Vars)
		deleted, delErr := DeleteAuth(auth.AccessUuid, client)
		if delErr != nil || deleted == 0 { //if any goes wrong
			response.RespondWithError(w, http.StatusUnauthorized, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, admin.MessageResponse{Status: true, Message: "Đã đăng xuất"})
	}
}

// ChangePassword controller for change password of account
func ChangePassword(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("values").(middlewares.Vars)
		_, err := middlewares.FetchAuth(userRole.AccessUuid, client)
		if err != nil {
			response.RespondWithError(w, http.StatusUnauthorized, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
			return
		}
		req := admin.ChangeAccountRequest{}
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

func DeleteAuth(givenUuid string, client *redis.Client) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
