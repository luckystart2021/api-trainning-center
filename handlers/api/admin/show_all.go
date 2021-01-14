package admin

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
)

func ShowAllAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("values").(middlewares.Vars)
		_, err := FetchAuth(userRole.AccessUuid, client)
		if err != nil {
			response.RespondWithError(w, http.StatusUnauthorized, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
			return
		}
		if userRole.Role != ADMIN {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Bạn không có quyền truy cập"))
			return
		}
		showAllAccount, err := service.ShowAllAccount()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAllAccount)
	}
}
