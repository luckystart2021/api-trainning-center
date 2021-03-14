package seo

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSeo) DeleteSeoTags(id int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := deleteSeoByRequest(st.db, id)
	if err != nil {
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa thành công"
	} else {
		resp.Status = false
		resp.Message = "Xóa không thành công"
	}
	return resp, nil
}

func deleteSeoByRequest(db *sql.DB, id int) (int64, error) {
	query := `
	DELETE 
		FROM 
			article_tag
		WHERE 
			id=$1;
	`
	res, err := db.Exec(query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteSeoByRequest] delete seo tag in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] delete seo tag in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
