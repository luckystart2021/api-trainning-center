package schedule

import (
	"api-trainning-center/internal/models"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreSchedule) RetrieveSchedule(courseID int) (Schedule, error) {
	scheduleCourse := Schedule{}
	scheduleCourseLT := scheduleCourse.LyThuyet
	ctx := context.Background()
	scheduleR, err := models.Schedules(
		qm.Where("course_id = ?", courseID),
		qm.OrderBy("id"),
		qm.Load(models.ScheduleRels.ScheduleContents+"."+models.ScheduleContentRels.ScheduleSubjectScheduleSubjects),
	).All(ctx, st.db)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find schedule] error : ", err)
		return scheduleCourse, err
	}

	if len(scheduleR) == 0 {
		logrus.WithFields(logrus.Fields{}).Error("[Find schedule] error : ", err)
		return scheduleCourse, errors.New("Không có dữ liệu từ hệ thống")
	}

	for _, data := range scheduleR {
		scheduleResponse := ScheduleResponse{}
		scheduleResponse.SubjectName = data.SubjectName
		scheduleResponse.Teacher = data.TeacherName
		scheduleResponse.Time = data.Time
		scheduleResponse.TotalLT = data.TotalLythuyet
		scheduleResponse.TotalTH = data.TotalThuchanh
		if data.R == nil && len(data.R.ScheduleContents) == 0 {
			logrus.WithFields(logrus.Fields{}).Error("[Find ScheduleContents] error : ", err)
			return scheduleCourse, errors.New("Không có dữ liệu từ hệ thống")
		}
		for i, dataContent := range data.R.ScheduleContents {
			scheduleContent := Content{}
			scheduleContent.WeekDay = dataContent.Weekday
			scheduleContent.Date = dataContent.Date
			scheduleResponse.Schedule = append(scheduleResponse.Schedule, scheduleContent)
			if dataContent.R == nil && len(dataContent.R.ScheduleSubjectScheduleSubjects) == 0 {
				logrus.WithFields(logrus.Fields{}).Error("[Find ScheduleSubjectScheduleSubjects] error : ", err)
				return scheduleCourse, errors.New("Không có dữ liệu từ hệ thống")
			}
			for _, dataSubject := range dataContent.R.ScheduleSubjectScheduleSubjects {
				subjectContent := SubjectContent{}
				subjectContent.Name = dataSubject.Name
				subjectContent.Lt = dataSubject.LT
				subjectContent.Th = dataSubject.TH
				scheduleResponse.Schedule[i].SubjectContents = append(scheduleResponse.Schedule[i].SubjectContents, subjectContent)
			}
		}
		scheduleCourseLT = append(scheduleCourseLT, scheduleResponse)
	}
	scheduleCourse.LyThuyet = scheduleCourseLT
	// Get data cho thời khóa biểu học thực hành
	scheduleCourseTh := scheduleCourse.ThucHanh

	scheduleContentTH, err := models.ScheduleContents(
		models.ScheduleContentWhere.CourseID.IsNotNull(),
		models.ScheduleContentWhere.CourseID.EQ(null.IntFrom(courseID)),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find ScheduleContents] error : ", err)
		return scheduleCourse, err
	}

	if len(scheduleContentTH) == 0 {
		logrus.WithFields(logrus.Fields{}).Error("[Find ScheduleContents] error : ", err)
		return scheduleCourse, errors.New("Không có dữ liệu từ hệ thống")
	}

	for _, dataScheduleContent := range scheduleContentTH {
		thucHanhResponse := ThucHanhResponse{}
		thucHanhResponse.WeekDay = dataScheduleContent.Weekday
		thucHanhResponse.Date = dataScheduleContent.Date
		thucHanhResponse.SubjectName = dataScheduleContent.SubjectName.String
		thucHanhResponse.HourStudent = dataScheduleContent.HourStudent.String
		thucHanhResponse.KmStudent = dataScheduleContent.KMStudent.String
		thucHanhResponse.HourPerDateVehicle = dataScheduleContent.HourPerDateVehicle.Int
		thucHanhResponse.KmDateVehicle = dataScheduleContent.KMDateVehicle.Int
		scheduleCourseTh = append(scheduleCourseTh, thucHanhResponse)
	}
	scheduleCourse.ThucHanh = scheduleCourseTh
	return scheduleCourse, nil
}
