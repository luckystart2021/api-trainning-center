package schedule

import (
	"api-trainning-center/service/admin/schedule"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		st := schedule.NewStoreSchedule(db)
		router.Get("/schedule", SearchSchedule(st))
	}
}

func SearchSchedule(service schedule.IScheduleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		searchKey := r.URL.Query().Get("key")
		if searchKey == "" || len(searchKey) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập từ khoá tìm kiếm"))
			return
		}

		showResultByKey, err := service.SearchScheduleByCourseCode(searchKey)
		if err != nil {
			if err == sql.ErrNoRows {
				response.RespondWithError(w, http.StatusBadRequest, errors.New("Không có dữ liệu từ hệ thống"))
				return
			}
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		response.RespondWithJSON(w, http.StatusOK, showResultByKey)
	}
}
