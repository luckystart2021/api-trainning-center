package utils

import "api-trainning-center/service/constant"

func CalculatePageNumber(totalItemCount int) []PageNumberResponse {
	pages := []PageNumberResponse{}
	totalPage := (totalItemCount + constant.ItemsPerPage - 1) / constant.ItemsPerPage
	countPage := 1
	for i := 0; i < totalPage; i++ {
		page := PageNumberResponse{
			PageNumber: countPage,
		}
		countPage++
		pages = append(pages, page)
	}
	return pages
}

type PageNumberResponse struct {
	PageNumber int `json:"page_number"`
}
