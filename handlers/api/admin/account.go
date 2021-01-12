package admin

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/models/admin"
	"log"
	"strconv"

	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
)

const ADMIN = "ADMIN"

// CreateAccount zcontroller for creating new users
func CreateAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("values").(middlewares.Vars)
		_, err := FetchAuth(userRole.AccessUuid, client)
		if err != nil {
			response.RespondWithError(w, http.StatusUnauthorized, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
			return
		}
		log.Println("Role", userRole.Role)
		if userRole.Role != ADMIN {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Bạn không có quyền tạo tài khoản"))
			return
		}
		req := admin.AccountRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
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

func Login(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := admin.AccountRequest{}
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

func LogoutAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Context().Value("values").(middlewares.Vars)
		deleted, delErr := DeleteAuth(auth.AccessUuid, client)
		if delErr != nil || deleted == 0 { //if any goes wrong
			response.RespondWithError(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, "Successfully logged out")
	}
}

func DeleteAuth(givenUuid string, client *redis.Client) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func FetchAuth(givenUuid string, client *redis.Client) (uint64, error) {
	userid, err := client.Get(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func ChangePassword(service user.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := admin.ChangeAccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
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
