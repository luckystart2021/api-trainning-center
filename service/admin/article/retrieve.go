package article

import (
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Article struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
}

type Articles struct {
	Id              int64  `json:"id"`
	IdUser          int64  `json:"id_user"`
	IdChildCategory int64  `json:"id_child_category"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Detail          string `json:"detail"`
	Img             string `json:"img"`
	Meta            string `json:"meta"`
	Keyword         string `json:"keyword"`
	View            int64  `json:"view"`
	Status          bool   `json:"status"`
	IsDelete        bool   `json:"is_deleted"`
	CreatedAt       string `json:"created_at"`
	CreatedBy       string `json:"created_by"`
	UpdatedAt       string `json:"updated_at"`
	UpdatedBy       string `json:"updated_by"`
}

type ArticleDetail struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	Keyword     string `json:"keyword"`
	View        int64  `json:"view"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
}

var (
	statusInactive  = true
	isDeleteIsFalse = false
)

func (tc StoreArticle) ShowArticle(idArticle int, meta string) (ArticleDetail, error) {
	articleDetail := ArticleDetail{}
	article, err := retrieveArticle(tc.db, idArticle, statusInactive, isDeleteIsFalse, meta)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveArticle] error : ", err)
		return articleDetail, err
	}
	articleDetails := ArticleDetail{
		Title:       article.Title,
		Description: article.Description,
		Detail:      article.Detail,
		Img:         article.Img,
		Meta:        article.Meta,
		Keyword:     article.Keyword,
		View:        article.View,
		CreatedAt:   article.CreatedAt,
		CreatedBy:   article.CreatedBy,
	}
	return articleDetails, nil
}

func retrieveArticle(db *sql.DB, idArticle int, statusInactive, isDeleteIsFalse bool, metaS string) (Articles, error) {
	query := `
	select
		articles.id,
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
	where
		articles.id = $1
		and articles.status = $2
		and articles.is_deleted = $3
		and articles.meta = $4;				
	`
	rows := db.QueryRow(query, idArticle, statusInactive, isDeleteIsFalse, metaS)
	var id, view, idUser, idChildCategory int64
	var title, description, details, img, meta, keywordseo, createdBy, updateBy string
	var createdAt, updatedAt time.Time
	var status, isDeleted bool
	err := rows.Scan(&id, &idUser, &idChildCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticle] No Data  %v", err)
		return Articles{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticle] Scan error  %v", err)
		return Articles{}, err
	}
	article := Articles{
		Id:              id,
		IdUser:          idUser,
		IdChildCategory: idChildCategory,
		Title:           title,
		Description:     description,
		Detail:          details,
		Img:             "/files/img/news/" + img,
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
	return article, nil
}

func (tc StoreArticle) ShowArticles(idCategory int) ([]Article, error) {
	article := []Article{}
	articels, err := retrieveArticles(tc.db, idCategory, statusInactive, isDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowArticles] error : ", err)
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
	logrus.WithFields(logrus.Fields{}).Info("[ShowArticles] retrieve success")
	return article, nil
}

func retrieveArticles(db *sql.DB, idCategory int, statusInactive, isDeleteIsFalse bool) ([]Articles, error) {
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
		c2.id = $1
		and articles.status = $2
		and articles.is_deleted = $3
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, idCategory, statusInactive, isDeleteIsFalse)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] query error  %v", err)
		return articles, err
	}
	for rows.Next() {
		var idArticle, view, idUser, idChildCategory int64
		var title, description, details, img, meta, keywordseo, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		var status, isDeleted bool
		err = rows.Scan(&idArticle, &idUser, &idChildCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] Scan error  %v", err)
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
	return articles, nil
}
