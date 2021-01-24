package contact

import (
	"api-trainning-center/service/admin/contact"
	"api-trainning-center/service/response"
	"net/http"
)

func ShowContacts(service contact.IContactService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showContact, err := service.ShowContacts()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showContact)
	}
}
