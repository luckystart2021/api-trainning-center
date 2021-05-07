package holiday

import (
	"api-trainning-center/service/admin/holiday"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func getHolidays(service holiday.IHolidayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showHolidays, err := service.ShowHolidays()
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showHolidays)
	}
}

func getHoliday(service holiday.IHolidayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã ngày nghỉ không được rỗng"))
			return
		}

		idHoliday, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã ngày nghỉ không hợp lệ"))
			return
		}
		showHoliday, err := service.ShowHoliday(idHoliday)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showHoliday)
	}
}
