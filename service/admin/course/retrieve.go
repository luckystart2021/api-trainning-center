package course

import (
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Course struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	GraduationDate string `json:"graduation_date"`
	TestDate       string `json:"test_date"`
	TrainingSystem string `json:"training_system"`
	Status         bool   `json:"status"`
	Time           string `json:"time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	CreatedBy      string `json:"created_by"`
	UpdatedBy      string `json:"updated_by"`
}

var (
	statusActive   bool = true
	statusInActive bool = false
)

func (tc StoreCourse) ShowCoursesActive(systemID string) ([]Course, error) {
	course, err := RetrieveCourses(statusActive, tc.db, systemID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowCoursesActive] error : ", err)
		return []Course{}, err
	}

	return course, nil
}

func (tc StoreCourse) ShowCoursesInActive() ([]Course, error) {
	// course, err := RetrieveCourses(statusInActive, tc.db)
	// if err != nil {
	// 	logrus.WithFields(logrus.Fields{}).Error("[ShowCoursesInActive] error : ", err)
	// 	return []Course{}, err
	// }

	return []Course{}, nil
}

func (tc StoreCourse) ShowCourses(idCourse string) (Course, error) {
	course, err := RetrieveCourse(idCourse, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowCourses] error : ", err)
		return Course{}, err
	}

	return course, nil
}

func RetrieveCourse(idCourse string, db *sql.DB) (Course, error) {
	courses := Course{}
	query := `
	SELECT 
		id, code, name, start_date, end_date, graduation_date, test_date, 
		training_system, status, created_by, created_at, updated_by, updated_at, time
	FROM 
		course
	WHERE
		id = $1;`
	rows := db.QueryRow(query, idCourse)
	var graduationDate sql.NullTime
	var startDate, endDate, testDate, createdAt, updatedAt time.Time

	err := rows.Scan(&courses.Id, &courses.Code, &courses.Name, &startDate, &endDate, &graduationDate,
		&testDate, &courses.TrainingSystem, &courses.Status, &courses.CreatedBy, &createdAt, &courses.UpdatedBy, &updatedAt, &courses.Time)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[RetrieveCourse] No Data  %v", err)
		return courses, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] Scan error  %v", err)
		return courses, err
	}
	if graduationDate.Valid {
		courses.GraduationDate = utils.TimeIn(graduationDate.Time, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
	}
	courses.StartDate = utils.TimeIn(startDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
	courses.EndDate = utils.TimeIn(endDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
	courses.TestDate = utils.TimeIn(testDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
	courses.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	courses.UpdatedAt = utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)

	return courses, nil
}

func RetrieveCourses(status bool, db *sql.DB, systemID string) ([]Course, error) {
	courses := []Course{}
	query := `
	SELECT 
		id, code, name, start_date, end_date, graduation_date, test_date, 
		training_system, status, created_by, created_at, updated_by, updated_at, time
	FROM 
		course
	WHERE
		training_system = $1
	ORDER BY start_date DESC;
	`
	rows, err := db.Query(query, systemID)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] query error  %v", err)
		return courses, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		var err error
		var graduationDate sql.NullTime
		var id int
		var code, name, trainingSystem, createdBy, updatedBy string
		var startDate, endDate, testDate, createdAt, updatedAt time.Time
		var status bool
		var timeC string
		err = rows.Scan(&id, &code, &name, &startDate, &endDate,
			&graduationDate, &testDate, &trainingSystem, &status,
			&createdBy, &createdAt, &updatedBy, &updatedAt, &timeC)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] Scan error  %v", err)
			return courses, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		course := Course{
			Id:             id,
			Code:           code,
			Name:           name,
			StartDate:      utils.TimeIn(startDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY),
			EndDate:        utils.TimeIn(endDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY),
			TestDate:       utils.TimeIn(testDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY),
			TrainingSystem: trainingSystem,
			Status:         status,
			Time:           timeC,
			CreatedAt:      utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			UpdatedAt:      utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
			CreatedBy:      createdBy,
			UpdatedBy:      updatedBy,
		}
		if graduationDate.Valid {
			course.GraduationDate = utils.TimeIn(graduationDate.Time, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
		}
		courses = append(courses, course)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(courses) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] No Data  %v", err)
		return courses, errors.New("Không có dữ liệu từ hệ thống")
	}
	return courses, nil
}
