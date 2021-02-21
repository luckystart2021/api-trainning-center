package slide

import (
	"api-trainning-center/service/admin/slide"
	"api-trainning-center/service/response"
	"net/http"
)

func GetSlides(service slide.ISlideService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showSlides, err := service.ShowSlides()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showSlides)
	}
}
