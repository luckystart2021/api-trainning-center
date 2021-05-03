package child_subject

import (
	"api-trainning-center/internal/models"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreChildSubject) ShowChildSubjects(idChildSubject int) (models.ChildSubjectSlice, error) {
	ctx := context.Background()
	childSubjects, err := models.ChildSubjects(
		qm.Where("subject_id = ?", idChildSubject),
		qm.OrderBy("id"),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindAllChildSubjects] error : ", err)
		return nil, err
	}
	if childSubjects == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return childSubjects, nil
}

func (st StoreChildSubject) ShowChildSubject(childSubjectID int) (models.ChildSubject, error) {
	subjectR := models.ChildSubject{}
	ctx := context.Background()
	childSubject, err := models.FindChildSubject(ctx, st.db, childSubjectID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindChildSubject] error : ", err)
		return subjectR, err
	}
	if childSubject == nil {
		return subjectR, errors.New("Không có dữ liệu từ hệ thống")
	}
	return *childSubject, nil
}
