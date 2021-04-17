package course

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreCourse) ActiveCourse(idCourse int, userName string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	ctx := context.Background()
	course, err := models.FindCourse(ctx, st.db, idCourse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[ActiveCourse] Update Active Course DB err  %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	course.Status = true
	course.UpdatedBy = userName
	rowsAff, err := course.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ActiveCourse] Update Active Course error : ", err)
		return response, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		response.Status = true
		response.Message = "Đã kích hoạt khóa học thành công"
	} else {
		response.Status = false
		response.Message = "Kích hoạt khóa học không thành công"
	}
	return response, nil
}

func (st StoreCourse) InActiveCourse(idCourse int, userName string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	ctx := context.Background()
	course, err := models.FindCourse(ctx, st.db, idCourse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[InActiveCourse] Update InActive Course DB err  %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	course.Status = false
	course.UpdatedBy = userName
	rowsAff, err := course.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[InActiveCourse] Update InActive Course error : ", err)
		return response, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		response.Status = true
		response.Message = "Hủy kích hoạt khóa học thành công"
	} else {
		response.Status = false
		response.Message = "Hủy kích hoạt khóa học không thành công"
	}
	return response, nil
}
