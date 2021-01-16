package middlewares

import (
	"api-trainning-center/service/response"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Vars struct {
	Role       string
	AccessUuid string
	UserName   string
}

// AuthJwtVerify verify token and add userID to the request context
func AuthJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := ExtractToken(r)
		if err != nil {
			log.Println("tokenString err ", err)
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã xác thực không tồn tại, vui lòng đăng nhập lại"))
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil {
			log.Print("Parse token error ", err)
			response.RespondWithError(w, http.StatusForbidden, errors.New("Phiên đăng nhập đã hết hạn, vui lòng đăng nhập lại"))
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response.RespondWithError(w, http.StatusForbidden, errors.New("Mã xác thực không hợp lệ"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || claims["Access_uuid"] == "" || claims["Role"] == "" || claims["ExpiresAt"] == "" || claims["UserID"] == "" {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Mã xác thực không hợp lệ: xác thực không thành công"))
			return
		}

		accessUuid, ok := claims["Access_uuid"].(string)
		if !ok {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Mã xác thực không hợp lệ: xác thực không thành công"))
			return
		}

		role, ok := claims["Role"].(string)
		if !ok {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Mã xác thực không hợp lệ: xác thực không thành công"))
			return
		}

		userName, ok := claims["UserID"].(string)
		if !ok {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Mã xác thực không hợp lệ: xác thực không thành công"))
			return
		}

		ctx := context.WithValue(r.Context(), "values", Vars{Role: role, AccessUuid: accessUuid, UserName: userName}) // adding the Role to the context
		if count := getTokenRemainingValidity(claims["ExpiresAt"]); count == -1 {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Mã xác thực đã hết hạn, vui lòng đăng nhập lại"))
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return -1
}

func ExtractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		return "", errors.New("Mã thông báo không được cung cấp hoặc không đúng định dạng")
	}
	return strArr[1], nil
}
