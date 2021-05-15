package course

import (
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type SearchRequest struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func Router(db *sql.DB) func(chi.Router) {
	st := course.NewStoreCourse(db)
	return func(router chi.Router) {
		router.Post("/course/filter", GetCourse(st))
	}
}

func GetCourse(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := SearchRequest{}
		err := json.NewDecoder(r.Body).Decode(&resp)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		startTime, err := utils.ParseStringToTime(resp.StartTime)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		endTime, err := utils.ParseStringToTime(resp.EndTime)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		respS, err := service.ShowCoursesByDate(startTime, endTime)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, respS)
	}
}
