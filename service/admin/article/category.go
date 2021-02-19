package article

import (
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Categories struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	IdCategory int64  `json:"id_category"`
	Meta       string `json:"meta"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedAt  string `json:"updated_at"`
	UpdatedBy  string `json:"updated_by"`
}

type CategoriesResponse struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

var (
	childCategoryIsDeleteIsFalse = false
	childCategoryIsDeleteIsTrue  = true
)

func (tc StoreArticle) ShowCategories(idCategoryParent int) ([]CategoriesResponse, error) {
	categoriesResponse := []CategoriesResponse{}
	categories, err := RetrieveCategories(tc.db, idCategoryParent, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowCategories] error : ", err)
		return categoriesResponse, err
	}
	for _, data := range categories {
		categoryResponse := CategoriesResponse{
			Id:    data.Id,
			Title: data.Title,
			Meta:  data.Meta,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse, nil
}

func RetrieveCategories(db *sql.DB, idCategoryParent int, childCategoryIsDeleteIsFalse bool) ([]Categories, error) {
	categories := []Categories{}
	query := `
	select
		id,
		title,
		id_category,
		meta,
		created_at,
		created_by,
		updated_at,
		updated_by
	from
		child_category cc
	where
		id_category = $1
		and is_deleted = $2
	order by
		id
	`
	rows, err := db.Query(query, idCategoryParent, childCategoryIsDeleteIsFalse)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategories] query error  %v", err)
		return categories, err
	}
	for rows.Next() {
		var id, idCategory int64
		var title, meta, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		err = rows.Scan(&id, &title, &idCategory, &meta, &createdAt, &createdBy, &updatedAt, &updateBy)

		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategories] Scan error  %v", err)
			return categories, err
		}
		category := Categories{
			Id:         id,
			Title:      title,
			IdCategory: idCategory,
			Meta:       meta,
			CreatedAt:  utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			CreatedBy:  createdBy,
			UpdatedAt:  utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			UpdatedBy:  updateBy,
		}
		categories = append(categories, category)
	}
	if len(categories) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategories] No Data  %v", err)
		return categories, errors.New("Không có dữ liệu từ hệ thống")
	}
	return categories, nil
}
