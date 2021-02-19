package article

import (
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

func (tc StoreArticle) ShowChildArticles(idChildCategoryP int, metaChild, metaParent string) ([]ChildCategoryNewsList, error) {
	childCategoryNewsList := []ChildCategoryNewsList{}
	childCategories, err := retrieveChildCategories(tc.db, idChildCategoryP, metaChild, metaParent, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowChildArticles] error : ", err)
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

func retrieveChildCategories(db *sql.DB, idChildCategoryP int, metaChild, metaParent string, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool) ([]Articles, error) {
	articles := []Articles{}
	query := `
	select
		articles.id ,
		articles.id_user,
		articles.id_child_category,
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
	from
		articles
	inner join child_category c on
		c.id = articles.id_child_category
	inner join category c1 on
		c.id_category = c1.id
	where
		c.id = $1
		and c.meta = $2
		and articles.status = $3
		and articles.is_deleted = $4
		and c1.meta = $5
		and c.is_deleted = $6
	order by
		articles.created_at desc;
	`
	rows, err := db.Query(query, idChildCategoryP, metaChild, statusActive, isDeleteIsFalse, metaParent, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] query error  %v", err)
		return articles, err
	}
	for rows.Next() {
		var idArticle, view, idUser, idChildCategory int64
		var title, description, details, img, meta, keywordseo, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		var status, isDeleted bool
		err = rows.Scan(&idArticle, &idUser, &idChildCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] Scan error  %v", err)
			return articles, err
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
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveChildCategories] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}
