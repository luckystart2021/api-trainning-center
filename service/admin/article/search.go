package article

import (
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (tc StoreArticle) ShowResultNewsByKey(searchKey string) ([]Article, error) {
	article := []Article{}
	articels, err := retrieveResultNewsByKey(tc.db, searchKey, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowResultNewsByKey] error : ", err)
		return article, err
	}
	for _, data := range articels {
		articleR := Article{
			Id:          data.Id,
			Title:       data.Title,
			Img:         "/files/img/news/" + data.Img,
			Meta:        data.Meta,
			Description: data.Description,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		article = append(article, articleR)
	}
	logrus.WithFields(logrus.Fields{}).Info("[ShowResultNewsByKey] retrieve success")
	return article, nil
}

func retrieveResultNewsByKey(db *sql.DB, searchKey string, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool) ([]Articles, error) {
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
	inner join category c2 on
		c.id_category = c2.id
	where
		articles.status = $1
		and articles.is_deleted = $2
		and c.is_deleted = $3
		and articles.title ilike $4
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse, "%"+searchKey+"%")
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveResultNewsByKey] query error  %v", err)
		return articles, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	for rows.Next() {
		var idArticle, view, idUser, idChildCategory int64
		var title, description, details, img, meta, keywordseo, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		var status, isDeleted bool
		err = rows.Scan(&idArticle, &idUser, &idChildCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)

		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveResultNewsByKey] Scan error  %v", err)
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
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveResultNewsByKey] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}
