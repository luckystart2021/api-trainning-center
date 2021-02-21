package slide

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/slide"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
)

type SlideRequest struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

func CreateSlide(service slide.ISlideService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		imageName, err := utils.FileUpload(r, "slide")
		// here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		req := SlideRequest{
			Title: r.FormValue("title"),
			Img:   imageName,
		}

		if err := validate(&req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.CreateSlide(userRole.UserName, req.Title, req.Img)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validate(s *SlideRequest) error {
	if s.Title == "" || len(s.Title) == 0 {
		return errors.New("Vui lòng nhập tiêu đề")
	}
	if len(s.Title) > 2000 {
		return errors.New("Tiêu đề không hợp lệ")
	}

	if s.Img == "" || len(s.Img) == 0 {
		return errors.New("Vui lòng nhập chọn hình ảnh")
	}
	return nil
}
