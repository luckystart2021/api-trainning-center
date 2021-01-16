package account

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/models/admin/account"
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
)

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
		response.RespondWithJSON(w, http.StatusOK, account.MessageResponse{Status: true, Message: "Đã đăng xuất"})
	}
}

func DeleteAuth(givenUuid string, client *redis.Client) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
