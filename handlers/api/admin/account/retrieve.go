package account

import (
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func RetrieveAccounts(service user.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showAllAccount, err := service.ShowAllAccount()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAllAccount)
	}
}

func RetrieveAccount(service user.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		if username == "" || len(username) == 0 || len(username) > 50 {
			logrus.WithFields(logrus.Fields{}).Error("[RetrieveAccount] param is null")
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Tên đăng nhập không tồn tại"))
			return
		}
		logrus.WithFields(logrus.Fields{}).Info("[RetrieveAccount] param is ", username)
		showAccount, err := service.ShowAccount(username)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAccount)
	}
}
