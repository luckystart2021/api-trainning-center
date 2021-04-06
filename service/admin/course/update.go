package course

import (
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreCourse) UpdateCourse(idCourse int, userName, name, startDate, endDate, graduationDate, testDate, trainingSystem, time string) (response.MessageResponse, error) {
	response := response.MessageResponse{}

	startTime, err := utils.ParseStringToTime(startDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày bắt đầu khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse start time error %v", err)
		return response, err
	}

	endTime, err := utils.ParseStringToTime(endDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày kết thúc khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse end time error %v", err)
		return response, err
	}

	graduationTime, err := utils.ParseStringToTime(graduationDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày tốt nghiệp khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse graduation time error %v", err)
		return response, err
	}

	testTime, err := utils.ParseStringToTime(testDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày thi sác hạch khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse test time error %v", err)
		return response, err
	}

	if err := updateCourseByRequest(st.db, idCourse, userName, name, trainingSystem, startTime, endTime, graduationTime, testTime, time); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Cập nhật khóa học mới thành công"
	return response, nil
}

func updateCourseByRequest(db *sql.DB, idCourse int, userName, name, trainingSystem string, startDate, endDate, graduationDate, testDate time.Time, timeC string) error {
	timeUpdate := time.Now()
	query := `
	UPDATE
		course
	SET
		name = $2,
		start_date = $3,
		end_date = $4,
		graduation_date = $5,
		test_date = $6,
		training_system = $7,
		updated_by = $8,
		updated_at = $9,
		time = $10
	WHERE
		id = $1;
	`
	_, err := db.Exec(query, idCourse, name, startDate, endDate, graduationDate, testDate, trainingSystem, userName, timeUpdate, timeC)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateCourseByRequest] Update DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
