package teacher

import (
	"api-trainning-center/internal/models"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreTeacher) ShowTeachers() (models.TeacherSlice, error) {
	ctx := context.Background()
	teachers, err := models.Teachers().All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findAllTeacher] error : ", err)
		return nil, err
	}
	if teachers == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return teachers, nil
}

func (st StoreTeacher) ShowTeacher(idTeacher int) (models.Teacher, error) {
	ctx := context.Background()
	teacher, err := models.FindTeacher(ctx, st.db, idTeacher)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[findOneTeacher] No Data  %v", err)
		return models.Teacher{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findOneTeacher] error : ", err)
		return models.Teacher{}, err
	}

	return *teacher, nil
}
