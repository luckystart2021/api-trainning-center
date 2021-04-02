package article

import (
	"api-trainning-center/models/admin/account"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreArticle) CreateArticle(idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	idUser, err := account.RetrieveAccountByUserName(userName, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RetrieveAccountByUserName] get err  %v", err)
		return resp, err
	}

	if err := CreateArticleByRequest(tc.db, idUser.Id, idChildCategoryP, userName, title, description, details, meta, keyWordSEO, image); err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateQuestion]Insert Question DB err  %v", err)
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm bài viết thành công"
	return resp, nil
}

func CreateArticleByRequest(db *sql.DB, idUser int, idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) error {
	query := `
	INSERT INTO articles
	(user_id, child_category_id, title, description, details, image, meta, keywordseo, created_by, updated_by)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`
	_, err := db.Exec(query, idUser, idChildCategoryP, title, description, details, image, meta, keyWordSEO, userName, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateArticleByRequest]Insert Article DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
