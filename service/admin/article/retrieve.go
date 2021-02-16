package article

import (
	"api-trainning-center/models/admin/account"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Article struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Img       string `json:"img"`
	Meta      string `json:"meta"`
	View      int64  `json:"view"`
	CreatedAt string `json:"created_at"`
}

type Articles struct {
	Id          int64  `json:"id"`
	IdUser      int64  `json:"id_user"`
	IdCategory  int64  `json:"id_category"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	Keyword     string `json:"keyword"`
	View        int64  `json:"view"`
	Status      bool   `json:"status"`
	IsDelete    bool   `json:"is_deleted"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
}

type ArticleDetail struct {
	Id          int64  `json:"id"`
	UserName    string `json:"user_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	Keyword     string `json:"keyword"`
	View        int64  `json:"view"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	Tags        []Tag  `json:"tags"`
}

type Tag struct {
	Name string `json:"name"`
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
	tags, err := retrieveTags(tc.db, article.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveTags] error : ", err)
		return articleDetail, err
	}
	account, err := account.RetrieveAccountInActiveById(article.IdUser, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[RetrieveAccountInActiveById] error : ", err)
		return articleDetail, err
	}
	articleDetails := ArticleDetail{
		Title:    article.Title,
		UserName: account.UserName,
		Tags:     tags,
	}
	return articleDetails, nil
}

func retrieveTags(db *sql.DB, idArticle int64) ([]Tag, error) {
	tags := []Tag{}
	query := `
	select
		t.title
	from
		articles_tags at2
	inner join tags t on
		t.id = at2.id_tag
	where
		id_article = $1;
	`
	rows, err := db.Query(query, idArticle)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveTags] No Data  %v", err)
		return tags, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveTags] query error  %v", err)
		return tags, err
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveTags] Scan error  %v", err)
			return tags, err
		}
		tag := Tag{
			Name: name,
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func retrieveArticle(db *sql.DB, idArticle int, statusInactive, isDeleteIsFalse bool, metaS string) (Articles, error) {
	query := `
	select
		articles.id,
		articles.id_user,
		articles.id_category,
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
	left join category c on
		c.id = articles.id_category
	where
		articles.id = $1
		and articles.status = $2
		and articles.is_deleted = $3
		and articles.meta = $4
	group by
		articles.id
	order by
		articles.created_at desc;				
	`
	rows := db.QueryRow(query, idArticle, statusInactive, isDeleteIsFalse, metaS)
	var id, view, idUser, idCategory int64
	var title, description, details, img, meta, keywordseo, createdBy, updateBy string
	var createdAt, updatedAt time.Time
	var status, isDeleted bool
	err := rows.Scan(&id, &idUser, &idCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticle] No Data  %v", err)
		return Articles{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticle] Scan error  %v", err)
		return Articles{}, err
	}
	article := Articles{
		Id:          id,
		IdUser:      idUser,
		IdCategory:  idCategory,
		Title:       title,
		Description: description,
		Detail:      details,
		Img:         img,
		Meta:        meta,
		Keyword:     keywordseo,
		View:        view,
		Status:      status,
		IsDelete:    isDeleted,
		CreatedAt:   utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
		CreatedBy:   createdBy,
		UpdatedAt:   utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
		UpdatedBy:   updateBy,
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
			Id:        data.Id,
			Title:     data.Title,
			Img:       data.Img,
			Meta:      data.Meta,
			View:      data.View,
			CreatedAt: data.CreatedAt,
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
		articles.id_category,
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
	left join category c on
		c.id = articles.id_category
	where
		articles.id_category = $1
		and articles.status = $2
		and articles.is_deleted = $3
	group by
		articles.id
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, idCategory, statusInactive, isDeleteIsFalse)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] No Data  %v", err)
		return []Articles{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] query error  %v", err)
		return articles, err
	}
	for rows.Next() {
		var idArticle, view, idUser, idCategory int64
		var title, description, details, img, meta, keywordseo, createdBy, updateBy string
		var createdAt, updatedAt time.Time
		var status, isDeleted bool
		err = rows.Scan(&idArticle, &idUser, &idCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] Scan error  %v", err)
			return articles, err
		}
		article := Articles{
			Id:          idArticle,
			IdUser:      idUser,
			IdCategory:  idCategory,
			Title:       title,
			Description: description,
			Detail:      details,
			Img:         img,
			Meta:        meta,
			Keyword:     keywordseo,
			View:        view,
			Status:      status,
			IsDelete:    isDeleted,
			CreatedAt:   utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			CreatedBy:   createdBy,
			UpdatedAt:   utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			UpdatedBy:   updateBy,
		}
		articles = append(articles, article)
	}
	return articles, nil
}
