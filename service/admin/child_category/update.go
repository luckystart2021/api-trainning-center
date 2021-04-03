package child_category

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreChildCategory) UpdateChildCategory(id int, userName, title, meta string, idCategory int) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	if err := updateChildCategoryByRequest(st.db, userName, title, meta, idCategory, id); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Cập nhật danh mục thành công"
	return response, nil
}

func updateChildCategoryByRequest(db *sql.DB, userName, title, meta string, idCategory, id int) error {
	timeUpdate := time.Now()
	query := `
	UPDATE child_category
		SET 
			title=$1,
		    category_id=$2,
		    meta=$3, 
		    updated_at=$4, 
			updated_by=$5
	WHERE id=$6;
	`
	_, err := db.Exec(query, title, idCategory, meta, timeUpdate, userName, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateChildCategoryByRequest] update DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
