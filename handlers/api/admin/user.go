package admin

import (
	"api-trainning-center/models"
	"api-trainning-center/service/response"
	"api-trainning-center/service/user"
	"api-trainning-center/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

// CreateAccount zcontroller for creating new users
func CreateAccount(service user.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.AccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.Validate(""); err != nil {
			// If input is wrong, return an HTTP error
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}

		if _, err := req.IsValid(); err != nil {
			// Check Role have is valid
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}

		resp, err := service.CreateAccount(req)
		if err != nil {
			render.Render(w, r, utils.ServerErrorRenderer(err))
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func Login(service user.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.AccountRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.Validate("login"); err != nil {
			// If input is wrong, return an HTTP error
			response.ERROR(w, http.StatusBadRequest, err)
			return
		}

		// resp, err := service.Login(req)
		// if err != nil {
		// 	render.Render(w, r, utils.ServerErrorRenderer(err))
		// 	return
		// }

	}
}
