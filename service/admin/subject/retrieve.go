package subject

import (
	"api-trainning-center/internal/models"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSubject) ShowSubjects() (models.SubjectSlice, error) {
	ctx := context.Background()
	subjects, err := models.Subjects().All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findAllSubjects] error : ", err)
		return nil, err
	}
	if subjects == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return subjects, nil
}
