package register

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/admin/register"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"encoding/json"
	"errors"
	"net/http"
)

type RegisterRequest struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	ClassID      int    `json:"class_id"`
	TeacherID    int    `json:"teacher_id"`
	GroundNumber string `json:"ground_number"`
}

func Register(service register.IRegisterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := RegisterRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
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

		resp, err := service.CreateRegister(input)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (r RegisterRequest) validate() error {
	if r.StartDate == "" {
		return errors.New("Thời gian bắt đầu chưa được nhập")
	}

	if r.EndDate == "" {
		return errors.New("Thời gian kết thúc chưa được nhập")
	}

	if r.StartDate == r.EndDate {
		return errors.New("Thời gian bắt đầu và thời gian kết thúc không được trùng nhau")
	}

	if r.ClassID == 0 {
		return errors.New("Lớp học chưa được nhập")
	}

	if r.TeacherID == 0 {
		return errors.New("Giáo viên chưa được nhập")
	}

	return nil
}
