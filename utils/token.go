package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessUuid  string
	AccessToken string
	UserID      string
	Role        string
	AtExpires   int64
}

// EncodeAuthToken signs authentication token
func EncodeAuthToken(us string, role string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AccessUuid = uuid.NewV4().String()
	td.UserID = us
	td.Role = role
	td.AtExpires = time.Now().Add(time.Hour * 24).Unix()

	var err error
	// Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	claims := jwt.MapClaims{}
	claims["UserID"] = td.UserID
	claims["Role"] = td.Role
	claims["ExpiresAt"] = td.AtExpires
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	td.AccessToken, err = token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return td, err
	}
	return td, nil
}

func CreateAuth(userName string, td *TokenDetails, client *redis.Client) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, td.AccessToken, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}
