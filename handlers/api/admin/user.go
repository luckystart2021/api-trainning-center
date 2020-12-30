package admin

import (
	"api-trainning-center/models"
	"api-trainning-center/service/response"
	"api-trainning-center/service/user"
	"api-trainning-center/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

// CreateAccount handle Request
func CreateAccount(service user.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.AccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.Validate(); err != nil {
			// If input is wrong, return an HTTP error
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}

		if _, err := req.IsValid(); err != nil {
			// Check Role have is valid
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}

		hashPassword, err := HashPassword(req.PassWord)
		if err != nil {
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}
		req.PassWord = string(hashPassword)

		resp, err := service.CreateAccount(req)
		if err != nil {
			render.Render(w, r, utils.ServerErrorRenderer(err))
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

// HashPassword hashes password from user input
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 is the cost for hashing the password.
	if err != nil {
		return nil, errors.New("hashes password error")
	}
	return bytes, err
}

// CheckPasswordHash checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}
