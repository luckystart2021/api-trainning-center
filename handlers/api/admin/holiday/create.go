package holiday

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/admin/holiday"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"net/http"
)

func createHoliday(service holiday.IHolidayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Holiday{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := validateCreate(req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := service.CreateHoliday(req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateCreate(s models.Holiday) error {
	if len(s.Date) == 0 {
		return errors.New("Ngày nghỉ chưa được nhập")
	}
	if len(s.Date) > 10 || !validate.CheckDate(s.Date) {
		return errors.New("Ngày nghỉ không đúng định dạng")
	}

	if len(s.Name.String) == 0 {
		return errors.New("Nội dung ngày nghỉ chưa nhập")
	}

	_, err := utils.ParseStringToTime(s.Date)
	if err != nil {
		return errors.New("Ngày nghỉ không hợp lệ")
	}
	return nil
}
