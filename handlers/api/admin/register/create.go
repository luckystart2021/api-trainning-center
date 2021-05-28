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
