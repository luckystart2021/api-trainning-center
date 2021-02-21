package slide

import (
	"api-trainning-center/service/admin/slide"
	"api-trainning-center/service/response"
	"net/http"
)

type SlideResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Img       string `json:"img"`
	Hide      bool   `json:"hide"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"create_by"`
}

func ShowSlides(service slide.ISlideService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resps := []SlideResponse{}
		showSlides, err := service.ShowSlidesAdmin()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		for _, data := range showSlides {
			resp := SlideResponse{
				Id:        data.Id,
				Title:     data.Title,
				Img:       "/files/img/slide/" + data.Img,
				Hide:      data.Hide,
				CreatedAt: data.CreatedAt,
				CreatedBy: data.CreatedBy,
			}
			resps = append(resps, resp)
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resps)
	}
}
