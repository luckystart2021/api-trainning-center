package seo

import (
	modelSeo "api-trainning-center/models/admin/seo"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSeo) UpdateSeo(id int, req modelSeo.SeoRequest) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updateSeoByRequest(st.db, id, req)
	if err != nil {
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật seo thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật seo không thành công"
	}
	return resp, nil
}

func (st StoreSeo) UpdateSeoTags(id int, name string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updateSeoTagByRequest(st.db, id, name)
	if err != nil {
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật không thành công"
	}
	return resp, nil
}

func updateSeoTagByRequest(db *sql.DB, id int, name string) (int64, error) {
	query := `
	UPDATE
		article_tag
	SET
		name = $2
	WHERE
		id = $1;
	`
	res, err := db.Exec(query, id, name)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateSeoTagByRequest] update seo tag in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update seo tag in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}

func updateSeoByRequest(db *sql.DB, id int, req modelSeo.SeoRequest) (int64, error) {
	query := `
	UPDATE
		seo
	SET
		description = $2,
		keywords = $3,
		fb_app_id = $4,
		og_title = $5,
		og_url = $6,
		og_image = $7,
		og_description = $8,
		og_site_name = $9,
		og_see_also = $10,
		og_locale = $11,
		article_author = $12,
		twitter_card = $13,
		twitter_url = $14,
		twitter_title = $15,
		twitter_description = $16,
		twitter_image = $17,
		author = $18,
		generator = $19,
		copyright = $20
	WHERE
		id = $1;
	`
	res, err := db.Exec(query, id, req.Description, req.Keywords, req.FbAppId, req.OgTitle, req.OgUrl, req.OgImage, req.OgDescription,
		req.OgSiteName, req.OgSeeAlso, req.OgLocale, req.ArticleAuthor, req.TwitterCard, req.TwitterUrl, req.TwitterTitle,
		req.TwitterDescription, req.TwitterImage, req.Author, req.Generator, req.Copyright)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateSeoByRequest] update seo in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update seo in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
