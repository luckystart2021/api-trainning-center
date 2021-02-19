package child_category

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	childCategoryIsDeleteIsFalse = false
	childCategoryIsDeleteIsTrue  = true
)

func (st StoreChildCategory) ShowChildCategories(idCategoryParent int) ([]Categories, error) {
	categories, err := retrieveCategories(st.db, idCategoryParent)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowChildCategories] error : ", err)
		return []Categories{}, err
	}
	return categories, nil
}

func (st StoreChildCategory) ShowChildCategory(idChildCategory int) (article.Categories, error) {
	categories, err := retrieveCategory(st.db, idChildCategory)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowChildCategory] error : ", err)
		return article.Categories{}, err
	}
	return categories, nil
}

func retrieveCategory(db *sql.DB, idChildCategory int) (article.Categories, error) {
	category := article.Categories{}
	query := `
	SELECT 
		id, title, id_category, meta, created_at, created_by, updated_at, updated_by
	FROM 
		child_category
	WHERE id = $1;
	`
	rows := db.QueryRow(query, idChildCategory)
	var id, idCategory int64
	var title, meta, createdBy, updateBy string
	var createdAt, updatedAt time.Time
	err := rows.Scan(&id, &title, &idCategory, &meta, &createdAt, &createdBy, &updatedAt, &updateBy)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategory] No Data  %v", err)
		return category, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategory] Scan error  %v", err)
		return category, errors.New("Lỗi hệ thống")
	}
	categoryR := article.Categories{
		Id:         id,
		Title:      title,
		IdCategory: idCategory,
		Meta:       meta,
		CreatedAt:  utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
		CreatedBy:  createdBy,
		UpdatedAt:  utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
		UpdatedBy:  updateBy,
	}
	return categoryR, nil
}

type Categories struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	IdCategory int64  `json:"id_category"`
	Meta       string `json:"meta"`
	IsDeleted  bool   `json:"is_deleted"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedAt  string `json:"updated_at"`
	UpdatedBy  string `json:"updated_by"`
}

func retrieveCategories(db *sql.DB, idCategoryParent int) ([]Categories, error) {
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
		updated_by,
		is_deleted
	from
		child_category cc
	where
		id_category = $1
	order by
		id
	`
	rows, err := db.Query(query, idCategoryParent)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategories] query error  %v", err)
		return categories, err
	}
	for rows.Next() {
		var id, idCategory int64
		var title, meta, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		var isDeleted bool
		err = rows.Scan(&id, &title, &idCategory, &meta, &createdAt, &createdBy, &updatedAt, &updateBy, &isDeleted)

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
			IsDeleted:  isDeleted,
		}
		categories = append(categories, category)
	}
	if len(categories) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCategories] No Data  %v", err)
		return categories, errors.New("Không có dữ liệu từ hệ thống")
	}
	return categories, nil
}
