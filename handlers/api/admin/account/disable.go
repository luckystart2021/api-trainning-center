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

func DisableAccount(service user.IUserService, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		if username == "" || len(username) == 0 || len(username) > 50 {
			logrus.WithFields(logrus.Fields{}).Error("[DisableAccount] param is null")
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Tên đăng nhập không tồn tại"))
			return
		}
		logrus.WithFields(logrus.Fields{}).Info("[DisableAccount] param is ", username)
		deleteAccount, err := service.DisableAccountByUserName(username)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, deleteAccount)
	}
}
