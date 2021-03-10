package seo

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSeo) CreateSeoTag(name string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := createSeoTagByRequest(st.db, name); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm thẻ thành công"
	return resp, nil
}

func createSeoTagByRequest(db *sql.DB, name string) error {
	query := `
	INSERT INTO article_tag("name")
	VALUES($1);
	`
	_, err := db.Exec(query, name)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[createSeoTagByRequest]Insert seo tag DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
