package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"net/http"
)

func NotificationNews(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notificationNews, err := service.GetNotificationNews()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, notificationNews)
	}
}
