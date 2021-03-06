package slide

import (
	"api-trainning-center/service/admin/slide"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type SlideRequestUpdate struct {
	Title string
	Img   string
	Hide  string
}

func UpdateSlide(service slide.ISlideService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không được rỗng"))
			return
		}

		idSlide, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không hợp lệ"))
			return
		}

		imageName, err := utils.FileUpload(r, "slide")
		//here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
			//checking whether any error occurred retrieving image
		}
		req := SlideRequestUpdate{
			Title: r.FormValue("title"),
			Hide:  r.FormValue("hide"),
		}
		if imageName != "" || len(imageName) > 0 {
			req.Img = imageName
		}

		if err := validateUpdate(&req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := service.UpdateSlide(idSlide, req.Title, req.Img, req.Hide)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateUpdate(s *SlideRequestUpdate) error {
	if s.Title == "" || len(s.Title) == 0 {
		return errors.New("Vui lòng nhập tiêu đề")
	}
	if len(s.Title) > 2000 {
		return errors.New("Tiêu đề không hợp lệ")
	}

	_, err := strconv.ParseBool(s.Hide)
	if err != nil {
		return errors.New("Ẩn/hiện slide không đúng định dạng")
	}

	return nil
}
