package admin

import (
	"api-trainning-center/service/admin"
	"net/http"
)

func createAccount(service admin.IUserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// req := models.AccountRequest{}
		// err := json.NewEncoder(r.Body).Decode(&req)
		// if err != nil {
		// 	return
		// }
		// response, err := service.CreateAccount(&req)
		// if err != nil {
		// 	render.Render(w, r, handlers.ServerErrorRenderer(err))
		// 	return
		// }
		// // send Result response
		// response.RespondWithJSON(w, http.StatusOK, response)
	}
}
