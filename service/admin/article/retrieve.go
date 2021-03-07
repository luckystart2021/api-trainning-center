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

type NotificationNews struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Meta      string `json:"meta"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
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

type AdminArticlesList struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	View        int64  `json:"view"`
	Status      bool   `json:"status"`
	IsDelete    bool   `json:"is_deleted"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
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
	statusActive    = true
	statusInActive  = false
	isDeleteIsFalse = false
	isDeleteIsTrue  = true
)

func (tc StoreArticle) ShowArticle(idArticle int, meta string) (ArticleDetail, error) {
	articleDetail := ArticleDetail{}
	article, err := retrieveArticle(tc.db, idArticle, statusActive, isDeleteIsFalse, meta)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveArticle] error : ", err)
		return articleDetail, err
	}
	countView, err := updateViewArticle(tc.db, idArticle, article.View)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[updateViewArticle] error : ", err)
		return articleDetail, err
	}
	articleDetails := ArticleDetail{
		Title:       article.Title,
		Description: article.Description,
		Detail:      article.Detail,
		Img:         article.Img,
		Meta:        article.Meta,
		Keyword:     article.Keyword,
		View:        countView,
		CreatedAt:   article.CreatedAt,
		CreatedBy:   article.CreatedBy,
	}
	return articleDetails, nil
}

func updateViewArticle(db *sql.DB, idArticle int, view int64) (int64, error) {
	countview := view + 1
	query := `
	update
		articles
	set
		view = $2
	where
		id = $1
	`
	_, err := db.Exec(query, idArticle, countview)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateViewArticle] Update view Article DB err  %v", err)
		return view, err
	}
	return countview, nil
}

func retrieveArticle(db *sql.DB, idArticle int, statusActive, isDeleteIsFalse bool, metaS string) (Articles, error) {
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
	rows := db.QueryRow(query, idArticle, statusActive, isDeleteIsFalse, metaS)
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
	articels, err := retrieveArticles(tc.db, idCategory, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
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

func retrieveArticles(db *sql.DB, idCategory int, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool) ([]Articles, error) {
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
		and c.is_deleted = $4
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, idCategory, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func retrieveArticlesByChildCategory(db *sql.DB, idCategory int, statusActive, isDeleteIsFalse bool) ([]Articles, error) {
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
	where
		c.id = $1
		and articles.status = $2
		and articles.is_deleted = $3
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, idCategory, statusActive, isDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesByChildCategory] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesByChildCategory] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesByChildCategory] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesByChildCategory] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func (tc StoreArticle) ShowArticlesByChildCategory(idChildCategory int) ([]AdminArticlesList, error) {
	articleLst := []AdminArticlesList{}
	articels, err := retrieveArticlesByChildCategory(tc.db, idChildCategory, statusActive, isDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowArticlesByChildCategory] error : ", err)
		return articleLst, err
	}
	for _, data := range articels {
		articleR := AdminArticlesList{
			Id:          data.Id,
			Title:       data.Title,
			Description: data.Description,
			Img:         "/files/img/news/" + data.Img,
			View:        data.View,
			Status:      data.Status,
			IsDelete:    data.IsDelete,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		articleLst = append(articleLst, articleR)
	}

	return articleLst, nil
}

func (tc StoreArticle) ShowArticlesDeleteByChildCategory() ([]AdminArticlesList, error) {
	articleLst := []AdminArticlesList{}
	articels, err := retrieveArticlesDeteledByChildCategory(tc.db, isDeleteIsTrue)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowArticlesByChildCategory] error : ", err)
		return articleLst, err
	}
	for _, data := range articels {
		articleR := AdminArticlesList{
			Id:          data.Id,
			Title:       data.Title,
			Description: data.Description,
			Img:         "/files/img/news/" + data.Img,
			View:        data.View,
			Status:      data.Status,
			IsDelete:    data.IsDelete,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		articleLst = append(articleLst, articleR)
	}

	return articleLst, nil
}

func (tc StoreArticle) ShowArticlesUnApproval() ([]AdminArticlesList, error) {
	articleLst := []AdminArticlesList{}
	articels, err := retrieveArticlesUnApproval(tc.db, statusInActive)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveArticlesUnApproval] error : ", err)
		return articleLst, err
	}
	for _, data := range articels {
		articleR := AdminArticlesList{
			Id:          data.Id,
			Title:       data.Title,
			Description: data.Description,
			Img:         "/files/img/news/" + data.Img,
			View:        data.View,
			Status:      data.Status,
			IsDelete:    data.IsDelete,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		articleLst = append(articleLst, articleR)
	}

	return articleLst, nil
}
func (tc StoreArticle) GetAllNews() ([]AdminArticlesList, error) {
	articleLst := []AdminArticlesList{}
	articels, err := retrieveAllArticles(tc.db, statusInActive)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveAllArticles] error : ", err)
		return articleLst, err
	}

	for _, data := range articels {
		articleR := AdminArticlesList{
			Id:          data.Id,
			Title:       data.Title,
			Description: data.Description,
			Img:         "/files/img/news/" + data.Img,
			View:        data.View,
			Status:      data.Status,
			IsDelete:    data.IsDelete,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		articleLst = append(articleLst, articleR)
	}

	return articleLst, nil
}

func retrieveAllArticles(db *sql.DB, statusInActive bool) ([]Articles, error) {
	articles := []Articles{}
	query := `
	SELECT
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
	FROM
		articles
	WHERE
		 articles.status = $1
		 AND articles.is_deleted = $2
	ORDER BY
		articles.created_at DESC;				
	`
	rows, err := db.Query(query, statusActive, isDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAllArticles] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAllArticles] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAllArticles] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAllArticles] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func retrieveArticlesUnApproval(db *sql.DB, statusInActive bool) ([]Articles, error) {
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
	where
		 articles.status = $1
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, statusInActive)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesUnApproval] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesUnApproval] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesUnApproval] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesUnApproval] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func retrieveArticlesDeteledByChildCategory(db *sql.DB, isDeleteIsTrue bool) ([]Articles, error) {
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
	where
		 articles.is_deleted = $1
	order by
		articles.created_at desc;				
	`
	rows, err := db.Query(query, isDeleteIsTrue)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesDeteledByChildCategory] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticles] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesDeteledByChildCategory] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticlesDeteledByChildCategory] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func (tc StoreArticle) ShowArticleById(idArticle int) (Articles, error) {
	articleInAdmin, err := retrieveArticleInAdmin(tc.db, idArticle)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowArticleById] error : ", err)
		return Articles{}, err
	}
	return articleInAdmin, nil
}

func retrieveArticleInAdmin(db *sql.DB, idArticle int) (Articles, error) {
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
	`
	rows := db.QueryRow(query, idArticle)
	var id, view, idUser, idChildCategory int64
	var title, description, details, img, meta, keywordseo, createdBy, updateBy string
	var createdAt, updatedAt time.Time
	var status, isDeleted bool
	err := rows.Scan(&id, &idUser, &idChildCategory, &title, &description, &details, &img, &meta, &keywordseo, &view, &status, &isDeleted, &createdAt, &createdBy, &updatedAt, &updateBy)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticleInAdmin] No Data  %v", err)
		return Articles{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveArticleInAdmin] Scan error  %v", err)
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

func (tc StoreArticle) ShowNews() ([]Article, error) {
	article := []Article{}
	articels, err := retrieveNews(tc.db, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowNews] error : ", err)
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
	logrus.WithFields(logrus.Fields{}).Info("[ShowNews] retrieve success")
	return article, nil
}

func retrieveNews(db *sql.DB, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool) ([]Articles, error) {
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
	order by
		articles.created_at desc
	limit 3;				
	`
	rows, err := db.Query(query, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNews] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNews] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNews] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNews] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func (tc StoreArticle) ShowFavoriteNews() ([]Article, error) {
	article := []Article{}
	articels, err := retrieveFavoriteNews(tc.db, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowFavoriteNews] error : ", err)
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
	logrus.WithFields(logrus.Fields{}).Info("[ShowFavoriteNews] retrieve success")
	return article, nil
}

func retrieveFavoriteNews(db *sql.DB, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool) ([]Articles, error) {
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
	order by
		articles.view desc
	limit 4;				
	`
	rows, err := db.Query(query, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveFavoriteNews] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveFavoriteNews] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveFavoriteNews] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveFavoriteNews] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func (tc StoreArticle) GetNotificationNews() ([]NotificationNews, error) {
	article := []NotificationNews{}
	articels, err := retrieveNotificationNew(tc.db, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[GetNotificationNews] error : ", err)
		return article, err
	}
	for _, data := range articels {
		articleR := NotificationNews{
			Id:        data.Id,
			Title:     data.Title,
			Meta:      data.Meta,
			CreatedAt: data.CreatedAt,
			CreatedBy: data.CreatedBy,
		}
		article = append(article, articleR)
	}
	logrus.WithFields(logrus.Fields{}).Info("[GetNotificationNews] retrieve success")
	return article, nil
}

func retrieveNotificationNew(db *sql.DB, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse bool) ([]Articles, error) {
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
		and c.id = 2
	order by
		articles.created_at desc
	limit 3;				
	`
	rows, err := db.Query(query, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNotificationNew] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNotificationNew] Scan error  %v", err)
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
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNotificationNew] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if len(articles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNotificationNew] No Data  %v", err)
		return articles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return articles, nil
}

func (tc StoreArticle) ShowArticlesHomePage(idCategory int) ([]Article, error) {
	article := []Article{}
	articels, err := retrieveArticles(tc.db, idCategory, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
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
