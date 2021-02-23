package suite

import (
	"api-trainning-center/service/admin/suite"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func GenarateTest(service suite.ISuiteTestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		number := chi.URLParam(r, "number_suite")
		if number == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã đề không không được rỗng"))
			return
		}
		rank := chi.URLParam(r, "rank")
		if rank == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã đề không không được rỗng"))
			return
		}
		genarateSuiteTest, err := service.GenarateSuiteTest(number, rank)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, genarateSuiteTest)
	}
}
