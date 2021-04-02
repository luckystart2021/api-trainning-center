package article

import (
	"api-trainning-center/models/admin/account"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (tc StoreArticle) UpdateArticle(idArticle, idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	idUser, err := account.RetrieveAccountByUserName(userName, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateArticle] get err  %v", err)
		return resp, err
	}

	count, err := updateArticleByRequest(tc.db, idArticle, idUser.Id, idChildCategoryP, userName, title, description, details, meta, keyWordSEO, image)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateArticle]Update Article DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật bài viết thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật bài viết không thành công"
	}

	return resp, nil
}

func updateArticleByRequest(db *sql.DB, idArticle, idUser, idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (int64, error) {
	timeUpdate := time.Now()
	var rowsAffected int64
	if image == "" || len(image) == 0 {
		query := `
	update
		articles
	set
		user_id = $2,
		child_category_id = $3,
		title = $4,
		description = $5,
		details = $6,
		meta = $7,
		keywordseo = $8,
		updated_by = $9,
		updated_at = $10
	where
		id = $1
	`
		res, err := db.Exec(query, idArticle, idUser, idChildCategoryP, title, description, details, meta, keyWordSEO, userName, timeUpdate)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateArticleByRequest] Update Article DB err  %v", err)
			return 0, err
		}
		// check how many rows affected
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateArticleByRequest] Update Article DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	} else {
		query1 := `
		update
			articles
		set
			user_id = $2,
			child_category_id = $3,
			title = $4,
			description = $5,
			details = $6,
			image = $7,
			meta = $8,
			keywordseo = $9,
			updated_by = $10,
			updated_at = $11
		where
			id = $1
		`
		res, err := db.Exec(query1, idArticle, idUser, idChildCategoryP, title, description, details, image, meta, keyWordSEO, userName, timeUpdate)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateArticleByRequest] Update Article DB err  %v", err)
			return 0, err
		}
		// check how many rows affected
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateArticleByRequest] Update Article DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}

	return rowsAffected, nil
}
