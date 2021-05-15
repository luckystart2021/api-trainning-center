package student

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreStudent) DeleteStudent(idStudent int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	student, err := models.FindStudent(ctx, st.db, idStudent)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[DeleteStudent] delete student in DB err  %v", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	rowsAff, err := student.Delete(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteStudent] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Xóa thành công"
	} else {
		resp.Status = false
		resp.Message = "Xóa không thành công"
	}

	return resp, nil
}
