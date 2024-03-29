package article

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (tc StoreArticle) DeleteArticleById(idArticle int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := deleteArticleById(tc.db, idArticle, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[DeleteArticleById]Delete Article DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa bài viết thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy bài viết"
	}
	return resp, nil
}

func (tc StoreArticle) UnDeleteArticleById(idArticle int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := unDeleteArticleById(tc.db, idArticle, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UnDeleteArticleById]UnDelete Article DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Phục hồi bài viết thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy bài viết"
	}
	return resp, nil
}

func unDeleteArticleById(db *sql.DB, idArticle int, userName string) (int64, error) {
	timeUpdate := time.Now()
	query := `
	update
		articles
	set
		is_deleted = $2,
		updated_at = $3,
		updated_by = $4
	where
		id = $1
	`
	res, err := db.Exec(query, idArticle, isDeleteIsFalse, timeUpdate, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[unDeleteArticleById] UnDelete Article DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] UnDelete Article DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}

func deleteArticleById(db *sql.DB, idArticle int, userName string) (int64, error) {
	timeUpdate := time.Now()
	query := `
	update
		articles
	set
		is_deleted = $2,
		updated_at = $3,
		updated_by = $4
	where
		id = $1
	`
	res, err := db.Exec(query, idArticle, isDeleteIsTrue, timeUpdate, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteArticleById] Delete Article DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] Delete Article DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
