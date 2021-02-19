package child_category

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	isDeleteIsTrue = true
)

func (tc StoreChildCategory) DeleteCategoryById(id int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := deleteCategoryById(tc.db, id, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[DeleteCategoryById]Delete child category DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa danh mục bài viết thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy danh mục bài viết"
	}
	return resp, nil
}

func deleteCategoryById(db *sql.DB, id int, userName string) (int64, error) {
	timeUpdate := time.Now()
	query := `
	update
		child_category
	set
		is_deleted = $2,
		updated_at = $3,
		updated_by = $4
	where
		id = $1
	`
	res, err := db.Exec(query, id, isDeleteIsTrue, timeUpdate, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteCategoryById] Delete ChildCategory DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] Delete ChildCategory DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
