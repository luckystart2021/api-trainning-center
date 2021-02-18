package child_category

import (
	"api-trainning-center/service/admin/article"

	"github.com/sirupsen/logrus"
)

func (st StoreChildCategory) ShowChildCategories(idCategoryParent int) ([]article.Categories, error) {
	categories, err := article.RetrieveCategories(st.db, idCategoryParent)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowChildCategories] error : ", err)
		return []article.Categories{}, err
	}
	return categories, nil
}
