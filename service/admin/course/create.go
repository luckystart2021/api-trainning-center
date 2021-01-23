package course

import (
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreCourse) CreateCourse(userName, name, startDate, endDate, graduationDate, testDate, trainingSystem string) (response.MessageResponse, error) {
	response := response.MessageResponse{}

	countCourse, err := CountCourse(st.db)
	if err != nil || countCourse < 0 {
		response.Status = false
		response.Message = "Tạo khóa học thất bại"
		logrus.WithFields(logrus.Fields{}).Errorf("[CountCourse] Count Course Error %v", err)
		return response, err
	}
	totalCourse := strconv.FormatInt(countCourse+1, 10)
	if err != nil {
		return response, err
	}
	
	codeCourse := trainingSystem + "-K" + totalCourse
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
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse start time error %v", err)
		return response, err
	}

	graduationTime, err := utils.ParseStringToTime(graduationDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày tốt nghiệp khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse start time error %v", err)
		return response, err
	}

	testTime, err := utils.ParseStringToTime(testDate)
	if err != nil {
		response.Status = false
		response.Message = "Ngày thi sác hạch khóa học không hợp lệ"
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStartTime] parse start time error %v", err)
		return response, err
	}

	if err := CreateCourseByRequest(st.db, userName, codeCourse, name, trainingSystem, startTime, endTime, graduationTime, testTime); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Tạo khóa học mới thành công"
	return response, nil
}

func CreateCourseByRequest(db *sql.DB, userName, codeCourse, name, trainingSystem string, startDate, endDate, graduationDate, testDate time.Time) error {
	query := `
	INSERT INTO course
		(code ,name , start_date, end_date, graduation_date, test_date, training_system, created_by, updated_by)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`
	_, err := db.Exec(query, codeCourse, name, startDate, endDate, graduationDate, testDate, trainingSystem, userName, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateCourseByRequest]Insert Course DB err  %v", err)
		return err
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
		return count, err
	}
	return count, nil
}
