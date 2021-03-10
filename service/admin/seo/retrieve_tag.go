package seo

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type SeoTagsResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (st StoreSeo) ShowSeoTags() ([]SeoTagsResponse, error) {
	seoTags, err := findAllSeoTag(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowSeoTags] error : ", err)
		return []SeoTagsResponse{}, err
	}
	return seoTags, nil
}

func (st StoreSeo) ShowDetailSeoTags(id int) (SeoTagsResponse, error) {
	detailSeoTag, err := retrieveDetailSeoTag(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowDetailSeoTags] error : ", err)
		return SeoTagsResponse{}, err
	}
	return detailSeoTag, nil
}

func retrieveDetailSeoTag(db *sql.DB, id int) (SeoTagsResponse, error) {
	seoTagsResponse := SeoTagsResponse{}
	query := `
	SELECT 
		id, name
	FROM article_tag
	WHERE id = $1;
	`
	rows := db.QueryRow(query, id)
	err := rows.Scan(&seoTagsResponse.Id, &seoTagsResponse.Name)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveDetailSeoTag] No Data  %v", err)
		return seoTagsResponse, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveDetailSeoTag] Scan error  %v", err)
	}
	return seoTagsResponse, nil
}

func findAllSeoTag(db *sql.DB) ([]SeoTagsResponse, error) {
	seoTags := []SeoTagsResponse{}
	query := `
	SELECT 
		id, name
	FROM article_tag;
	`
	rows, err := db.Query(query)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAllSeoTag] query error  %v", err)
		return seoTags, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		seoTag := SeoTagsResponse{}
		err = rows.Scan(&seoTag.Id, &seoTag.Name)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[findAllSeoTag] Scan error  %v", err)
			return seoTags, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		seoTags = append(seoTags, seoTag)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAllSeoTag] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(seoTags) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllSeoTag] No Data  %v", err)
		return seoTags, errors.New("Không có dữ liệu từ hệ thống")
	}
	return seoTags, nil

}
