package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// EncodeAuthToken signs authentication token
func EncodeAuthToken(us string, role string) (string, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	claims := jwt.MapClaims{}
	claims["userID"] = us
	claims["Role"] = role
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
