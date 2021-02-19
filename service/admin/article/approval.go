package article

import (
	"api-trainning-center/service/response"
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

func (tc StoreArticle) UnApprovalArticleById(idArticle int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := unApprovalArticleById(tc.db, idArticle, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UnApprovalArticleById] UnApprovalArticleById Article DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Không duyệt bài viết thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy bài viết"
	}
	return resp, nil
}

func (tc StoreArticle) ApprovalArticleById(idArticle int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := approvalArticleById(tc.db, idArticle, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[ApprovalArticleById] ApprovalArticleById Article DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Duyệt bài viết thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy bài viết"
	}
	return resp, nil
}

func approvalArticleById(db *sql.DB, idArticle int, userName string) (int64, error) {
	timeUpdate := time.Now()
	query := `
	update
		articles
	set
		status = $2,
		updated_at = $3,
		updated_by = $4
	where
		id = $1
	`
	res, err := db.Exec(query, idArticle, statusActive, timeUpdate, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[approvalArticleById] Approval Article DB err  %v", err)
		return 0, err
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] Approval Article DB err  %v", err)
		return 0, err
	}

	return rowsAffected, nil
}

func unApprovalArticleById(db *sql.DB, idArticle int, userName string) (int64, error) {
	timeUpdate := time.Now()
	query := `
	update
		articles
	set
		status = $2,
		updated_at = $3,
		updated_by = $4
	where
		id = $1
	`
	res, err := db.Exec(query, idArticle, statusInActive, timeUpdate, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[unApprovalArticleById] UnApproval Article DB err  %v", err)
		return 0, err
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[unApprovalArticleById] UnApproval Article DB err  %v", err)
		return 0, err
	}

	return rowsAffected, nil
}
