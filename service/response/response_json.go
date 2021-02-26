package response

import (
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UploadResponse struct {
	FileName string `json:"file_name"`
}

type ErrorResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		RespondWithJSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	RespondWithJSON(w, http.StatusBadRequest, nil)
}
