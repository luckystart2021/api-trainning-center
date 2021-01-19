package course

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

type Course struct {
	Id             int       `json:"id"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	GraduationDate time.Time `json:"graduation_date"`
	TestDate       time.Time `json:"test_date"`
	TrainingSystem string    `json:"training_system"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
}

type CourseStatus bool

const (
	ACTIVE   CourseStatus = true
	INACTIVE CourseStatus = false
)

func (tc StoreCourse) ShowCoursesActive() ([]Course, error) {
	course, err := retrieveCoursesActive(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowCoursesActive] error : ", err)
		return []Course{}, err
	}

	return course, nil
}

func retrieveCoursesActive(db *sql.DB) ([]Course, error) {
	courses := []Course{}
	query := `
	SELECT 
		id, code, name, start_date, end_date, graduation_date, test_date, 
		training_system, status, created_by, created_at, updated_by, updated_at
	FROM 
		course
	WHERE
		status = $1;`
	rows, err := db.Query(query, ACTIVE)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCoursesActive] query error  %v", err)
		return courses, err
	}
	for rows.Next() {
		var err error
		var graduationDate sql.NullTime
		var id int
		var code, name, trainingSystem, createdBy, updatedBy string
		var startDate, endDate, testDate, createdAt, updatedAt time.Time
		var status bool

		err = rows.Scan(&id, &code, &name, &startDate, &endDate, &graduationDate, &testDate, &trainingSystem, &status, &createdBy, &createdAt, &updatedBy, &updatedAt)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCoursesActive] Scan error  %v", err)
			return courses, err
		}
		course := Course{
			Id:             id,
			Code:           code,
			Name:           name,
			StartDate:      startDate,
			EndDate:        endDate,
			TestDate:       testDate,
			TrainingSystem: trainingSystem,
			Status:         status,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			CreatedBy:      createdBy,
			UpdatedBy:      updatedBy,
		}
		if graduationDate.Valid {
			course.GraduationDate = graduationDate.Time
		}
		courses = append(courses, course)
	}
	return courses, nil
}
