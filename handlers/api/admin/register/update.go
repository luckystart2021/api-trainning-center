package register

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/admin/register"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateRegister(service register.IRegisterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã đăng ký không được rỗng"))
			return
		}

		idRes, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã đăng ký không hợp lệ"))
			return
		}
		req := RegisterRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := req.validate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		input := models.RegisterGround{}
		startAt, err := utils.ParseStringToTimeRegister(req.StartDate, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Ngày không đúng định dạng"))
			return
		}

		endAt, err := utils.ParseStringToTimeRegister(req.EndDate, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Ngày không đúng định dạng"))
			return
		}

		if endAt.Before(startAt) {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Ngày kết thúc không được nhỏ hơn ngày bắt đầu"))
			return
		}
		input.StartDate = startAt
		input.EndDate = endAt
		input.ClassID = req.ClassID
		input.TeacherID = req.TeacherID
		input.GroundNumber = req.GroundNumber

		resp, err := service.UpdateRegister(idRes, input)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
