package child_category

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreChildCategory) CreateChildCategory(userName, title, meta string, idCategory int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := CreateChildCategoryByRequest(st.db, userName, title, meta, idCategory); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm danh mục thành công"
	return resp, nil
}

func CreateChildCategoryByRequest(db *sql.DB, userName, title, meta string, idCategory int) error {
	query := `
	INSERT INTO child_category
		(title, id_category, meta, created_by, updated_by)
	VALUES($1, $2, $3, $4, $5);
	`
	_, err := db.Exec(query, title, idCategory, meta, userName, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateChildCategoryByRequest]Insert child category DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
