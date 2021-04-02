package article

import (
	"api-trainning-center/service/constant"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type ChildCategoryNewsList struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
}

func (tc StoreArticle) CountChildArticles(metaChild, metaParent string) (int, error) {
	var count int
	query := `
	SELECT
		COUNT(*)
	FROM
		articles
	INNER JOIN child_category c ON
		c.id = articles.child_category_id
	INNER JOIN category c1 ON
		c.id_category = c1.id
	WHERE
		c.meta = $1
		AND articles.status = $2
		AND articles.is_deleted = $3
		AND c1.meta = $4
		AND c.is_deleted = $5
	`
	row := tc.db.QueryRow(query, metaChild, statusActive, isDeleteIsFalse, metaParent, childCategoryIsDeleteIsFalse)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CountChildArticles] error : ", err)
		return count, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return count, nil
}

func (tc StoreArticle) ShowChildArticles(metaChild, metaParent string, pageNo int) ([]ChildCategoryNewsList, error) {
	childCategoryNewsList := []ChildCategoryNewsList{}
	childCategories, err := retrieveChildCategories(tc.db, metaChild, metaParent, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse, pageNo)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveChildCategories] error : ", err)
		return childCategoryNewsList, err
	}
	for _, data := range childCategories {
		childCategoryNews := ChildCategoryNewsList{
			Id:          data.Id,
			Title:       data.Title,
			Img:         "/files/img/news/" + data.Img,
			Meta:        data.Meta,
			Description: data.Description,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		childCategoryNewsList = append(childCategoryNewsList, childCategoryNews)
	}

	return childCategoryNewsList, nil
}

func retrieveChildCategories(db *sql.DB, metaChild, metaParent string, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool, pageNo int) ([]Articles, error) {
	offset := (pageNo - 1) * constant.ItemsPerPage
	articles := []Articles{}
	query := `
	SELECT
		articles.id ,
		articles.user_id,
		articles.child_category_id,
		articles.title ,
		articles.description,
		articles.details ,
		articles.image ,
		articles.meta ,
		articles.keywordseo,
		articles.view,
		articles.status,
		articles.is_deleted,
		articles.created_at,
		articles.created_by,
		articles.updated_at,
		articles.updated_by
	FROM
		articles
	INNER JOIN child_category c ON
		c.id = articles.child_category_id
	INNER JOIN category c1 ON
		c.id_category = c1.id
	WHERE
		c.meta = $1
		and articles.status = $2
		and articles.is_deleted = $3
		and c1.meta = $4
		and c.is_deleted = $5
	ORDER BY
		articles.created_at DESC
	LIMIT $6 OFFSET $7;
	`
	rows, err := db.Query(query, metaChild, statusActive, isDeleteIsFalse, metaParent, childCategoryIsDeleteIsFalse, constant.ItemsPerPage, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] query error  %v", err)
		return articles, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		var idArticle, view, idUser, idChildCategory int64
		var title, description, details, img, meta, keywordseo, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		var status, isDeleted bool
		err = rows.Scan(&idArticle, &idUser, &idChildCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] Scan error  %v", err)
			return articles, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		article := Articles{
			Id:              idArticle,
			IdUser:          idUser,
			IdChildCategory: idChildCategory,
			Title:           title,
			Description:     description,
			Detail:          details,
			Img:             img,
			Meta:            meta,
			Keyword:         keywordseo,
			View:            view,
			Status:          status,
			IsDelete:        isDeleted,
			CreatedAt:       utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			CreatedBy:       createdBy,
			UpdatedAt:       utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			UpdatedBy:       updateBy,
		}
		articles = append(articles, article)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}
