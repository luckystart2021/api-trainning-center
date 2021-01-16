package account

import (
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// ChangePassword controller for change password of account
func DeleteAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		if username == "" || len(username) == 0 || len(username) > 50 {
			logrus.WithFields(logrus.Fields{}).Error("[DeleteAccount] param is null")
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Tên đăng nhập không tồn tại"))
			return
		}
		logrus.WithFields(logrus.Fields{}).Info("[DeleteAccount] param is ", username)
		deleteAccount, err := service.DeleteAccountByUserName(username)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, deleteAccount)
	}
}
