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

// SetContentTypeMiddleware sets content-type to json
func SetContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// AuthJwtVerify verify token and add userID to the request context
func AuthJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := ExtractToken(r)
		if tokenString == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Missing auth token"))
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		log.Print("Error ", err)
		if err != nil {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Invalid token, please login"))
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(r.Context(), "Role", claims["Role"]) // adding the Role to the context
		if count := getTokenRemainingValidity(claims["ExpiresAt"]); count == -1 {
			response.RespondWithError(w, http.StatusForbidden, errors.New("Sorry! your token expired!"))
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

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
