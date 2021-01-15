package middlewares

import (
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	VALUES = "values"
	ADMIN  = "ADMIN"
)

func CheckScopeAccess(client *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value(VALUES).(Vars)
			_, err := FetchAuth(userRole.AccessUuid, client)
			if err != nil {
				response.RespondWithError(w, http.StatusUnauthorized, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
				return
			}
			logrus.WithFields(logrus.Fields{}).Infof("User %s Role %s loging ", userRole.UserName, userRole.Role)
			if userRole.Role != ADMIN {
				response.RespondWithError(w, http.StatusBadRequest, errors.New("Bạn không có quyền truy cập"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func FetchAuth(givenUuid string, client *redis.Client) (uint64, error) {
	userid, err := client.Get(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}
