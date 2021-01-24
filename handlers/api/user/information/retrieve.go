package information

import (
	"api-trainning-center/service/admin/information"
	"api-trainning-center/service/response"
	"net/http"
)

func GetInformation(service information.IInformationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showInformation, err := service.ShowInformation()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showInformation)
	}
}
