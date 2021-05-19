package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreCost) UpdateCost(id int, req models.TrainingCost, userName string, courseID, classID int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	course, err := models.FindCourse(ctx, st.db, courseID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindCourse] FindCourse error : ", err)
		return resp, errors.New("Không tìm thấy khóa học")
	}
	if !course.Status {
		logrus.WithFields(logrus.Fields{}).Error("[FindCourse] FindCourse error : ", err)
		return resp, errors.New("Khóa học đã kết thúc")
	}

	class, err := models.FindClass(ctx, st.db, classID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindClass] FindClass error : ", err)
		return resp, errors.New("Không tìm thấy lớp học")
	}

	if class.IsDeleted {
		logrus.WithFields(logrus.Fields{}).Error("[FindClass] FindClass error : ", err)
		return resp, errors.New("Lớp học đã bị xóa")
	}

	cost, err := models.FindTrainingCost(ctx, st.db, int64(id))
	if cost == nil {
		return resp, errors.New("Thông tin không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTrainingCost] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	cost.Amount = req.Amount
	cost.Type = req.Type
	cost.Note = req.Note
	cost.UpdatedBy = userName
	cost.ClassID = classID
	cost.CourseID = courseID

	rowsAff, err := cost.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateCost] Update Cost error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin không thành công"
	}
	return resp, nil
}
