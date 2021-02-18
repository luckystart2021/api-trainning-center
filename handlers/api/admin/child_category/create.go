package child_category

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_category"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

type ChildCategoryRequest struct {
	Title      string `json:"title"`
	Meta       string `json:"meta"`
	IdCategory int    `json:"id_category"`
}

func CreateChildCategory(service child_category.IChildCategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := ChildCategoryRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := req.validate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.CreateChildCategory(userRole.UserName, req.Title, req.Meta, req.IdCategory)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c ChildCategoryRequest) validate() error {
	if c.Title == "" {
		return errors.New("Tên danh mục chưa được nhập")
	}
	if len(c.Title) > 1000 {
		return errors.New("Tên danh mục không hợp lệ")
	}
	if c.Meta == "" {
		return errors.New("Thẻ danh mục chưa được nhập")
	}
	if len(c.Meta) > 1000 {
		return errors.New("Thẻ danh mục không hợp lệ")
	}

	if c.IdCategory == 0 {
		return errors.New("Mã danh mục chưa được nhập")
	}

	return nil
}
