package schedule

import (
	"api-trainning-center/service/admin/schedule"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := schedule.NewStoreSchedule(db)
	return func(router chi.Router) {
		router.Get("/{id}/gen", GenerateSchedule(st))
	}
}

func GenerateSchedule(service schedule.IScheduleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không được rỗng"))
			return
		}

		idCourse, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không hợp lệ"))
			return
		}

		resp, err := service.GenerateSchedule(idCourse)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
