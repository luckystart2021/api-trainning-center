package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreCost) CreateCost(req models.TrainingCost, courseID, classID int, userName string) (response.MessageResponse, error) {
	ctx := context.Background()
	resp := response.MessageResponse{}

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

	req.CreatedBy = userName
	req.UpdatedBy = userName
	req.CourseID = courseID
	req.ClassID = classID
	err = req.Insert(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CreateCost] Create Cost error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	resp.Status = true
	resp.Message = "Thêm thành công"
	return resp, nil
}
