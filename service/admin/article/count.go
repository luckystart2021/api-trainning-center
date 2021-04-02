package article

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreArticle) CountArticles(idCategory int) (int, error) {
	var count int
	query := `
	SELECT 
		COUNT(*) 
	FROM
		articles
	INNER JOIN child_category c ON
		c.id = articles.child_category_id
	INNER JOIN category c2 ON
		c.id_category = c2.id
	WHERE
		c2.id = $1
		and articles.status = $2
		and articles.is_deleted = $3
		and c.is_deleted = $4
	`
	row := tc.db.QueryRow(query, idCategory, statusActive, isDeleteIsFalse, childCategoryIsDeleteIsFalse)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CountArticles] error : ", err)
		return count, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return count, nil
}
