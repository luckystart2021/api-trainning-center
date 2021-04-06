package course

import (
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreCourse) CreateCourse(userName, name, startDate, endDate, graduationDate, testDate, trainingSystem, timeC string) (response.MessageResponse, error) {
	response := response.MessageResponse{}

	startTime, err := utils.ParseStringToTime(startDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày bắt đầu khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse start time error %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	endTime, err := utils.ParseStringToTime(endDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày kết thúc khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse end time error %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	graduationTime, err := utils.ParseStringToTime(graduationDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày tốt nghiệp khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse graduation time error %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	testTime, err := utils.ParseStringToTime(testDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày thi sác hạch khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse test time error %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	if err := CreateCourseByRequest(st.db, userName, name, trainingSystem, startTime, endTime, graduationTime, testTime, timeC); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Tạo khóa học mới thành công"
	return response, nil
}

func CreateCourseByRequest(db *sql.DB, userName, name, trainingSystem string, startDate, endDate, graduationDate, testDate time.Time, time string) error {
	query := fmt.Sprintf(`
	INSERT INTO course
	(code, name , start_date, end_date, graduation_date, test_date, training_system, created_by, updated_by, time)
	(SELECT CONCAT('%s-K', COUNT(*)+1), $1, $2, $3, $4, $5, $6, $7, $7, $8 FROM course);
	`, trainingSystem)
	_, err := db.Exec(query, name, startDate, endDate, graduationDate, testDate, trainingSystem, userName, time)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateCourseByRequest]Insert Course DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}

func CountCourse(db *sql.DB) (int64, error) {
	var count int64
	query := `
	SELECT COUNT(*) FROM course;`
	row := db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CountCourse] count course query err  %v", err)
		return count, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return count, nil
}
