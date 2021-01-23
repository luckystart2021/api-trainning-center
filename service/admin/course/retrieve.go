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

var (
	statusActive   bool = true
	statusInActive bool = false
)

func (tc StoreCourse) ShowCoursesActive() ([]Course, error) {
	course, err := RetrieveCourses(statusActive, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowCoursesActive] error : ", err)
		return []Course{}, err
	}

	return course, nil
}

func (tc StoreCourse) ShowCourses(idCourse int) (Course, error) {
	course, err := RetrieveCourse(idCourse, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowCourses] error : ", err)
		return Course{}, err
	}

	return course, nil
}

func RetrieveCourse(idCourse int, db *sql.DB) (Course, error) {
	courses := Course{}
	query := `
	SELECT 
		id, code, name, start_date, end_date, graduation_date, test_date, 
		training_system, status, created_by, created_at, updated_by, updated_at
	FROM 
		course
	WHERE
		id = $1;`
	rows := db.QueryRow(query, idCourse)

	var graduationDate sql.NullTime
	err := rows.Scan(&courses.Id, &courses.Code, &courses.Name, &courses.StartDate, &courses.EndDate, &graduationDate,
		&courses.TestDate, &courses.TrainingSystem, &courses.Status, &courses.CreatedBy, &courses.CreatedAt, &courses.UpdatedBy, &courses.UpdatedAt)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] Scan error  %v", err)
		return courses, err
	}
	if graduationDate.Valid {
		courses.GraduationDate = graduationDate.Time
	}

	return courses, nil
}

func RetrieveCourses(status bool, db *sql.DB) ([]Course, error) {
	courses := []Course{}
	query := `
	SELECT 
		id, code, name, start_date, end_date, graduation_date, test_date, 
		training_system, status, created_by, created_at, updated_by, updated_at
	FROM 
		course
	WHERE
		status = $1;`
	rows, err := db.Query(query, status)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] query error  %v", err)
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
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveCourses] Scan error  %v", err)
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
