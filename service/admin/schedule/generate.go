package schedule

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Schedule struct {
	LyThuyet []ScheduleResponse `json:"ly_thuyet"`
	ThucHanh []ThucHanhResponse `json:"thuc_hanh"`
}

type ThucHanhResponse struct {
	WeekDay            string `json:"week_day"`
	Date               string `json:"date"`
	SubjectName        string `json:"subject_name"`
	HourStudent        string `json:"gio_hoc_vien"`
	KmStudent          string `json:"km_hoc_vien"`
	HourPerDateVehicle int    `json:"gio_ngay_xe"`
	KmDateVehicle      int    `json:"km_ngay_xe"`
}
type ScheduleResponse struct {
	SubjectName string    `json:"subject_name"`
	Time        int       `json:"time"`
	Teacher     string    `json:"teacher"`
	Schedule    []Content `json:"schedule"`
	TotalLT     int       `json:"total_lt"`
	TotalTH     int       `json:"total_th"`
}

type Content struct {
	WeekDay         string           `json:"week_day"`
	Date            string           `json:"date"`
	SubjectContents []SubjectContent `json:"subject_contents"`
}

type SubjectContent struct {
	Name string `json:"name"`
	Lt   int    `json:"lt"`
	Th   int    `json:"th"`
}

func ConvertDayString(day int) string {
	names := [...]string{
		"CN",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7"}
	// trả về tên của 1 hằng số Weekday từ mảng names bên trên
	return names[day]
}

var (
	holiday = map[string]string{
		"30-04-2020": "holiday",
		"01-05-2020": "holiday",
		"23-04-2020": "holiday",
		"10-05-2020": "holiday",
	}
	continueDate time.Time
)

func (st StoreSchedule) GenerateSchedule(courseId int) (Schedule, error) {
	scheduleResponses := Schedule{}
	course, _ := models.FindCourse(context.Background(), st.db, courseId)
	if course.TrainingSystem == "B2" {
		lyThuyet, err := generateB2(st.db, course)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[generateB2] error : ", err)
			return scheduleResponses, err
		}
		scheduleResponses.LyThuyet = lyThuyet
		thucHanh, err := generateThucHanhB2(st.db)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[generateB2] error : ", err)
			return scheduleResponses, err
		}
		scheduleResponses.ThucHanh = thucHanh
	}
	err := saveSchedules(courseId, st.db, scheduleResponses)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[saveSchedules] error : ", err)
		return scheduleResponses, err
	}
	return scheduleResponses, nil
}

func saveSchedules(courseId int, st *sql.DB, scheduleResponses Schedule) error {
	ctx := context.Background()
	lyThuyets := scheduleResponses.LyThuyet
	for _, lyThuyet := range lyThuyets {
		schedule := models.Schedule{}
		schedule.SubjectName = lyThuyet.SubjectName
		schedule.TeacherName = lyThuyet.Teacher
		schedule.Time = lyThuyet.Time
		err := schedule.Insert(ctx, st, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[Insert] Create schedule error : ", err)
			return errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		scheduleD, err := models.Schedules(
			qm.OrderBy("id DESC"),
			qm.Limit(1)).One(ctx, st)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[Schedules] Find Schedules error : ", err)
			return errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		for _, content := range lyThuyet.Schedule {
			contentSchedule := models.ScheduleContent{}
			contentSchedule.Weekday = content.WeekDay
			contentSchedule.Date = content.Date
			contentSchedule.ScheduleID = null.Int64From(scheduleD.ID)
			err := contentSchedule.Insert(ctx, st, boil.Infer())
			if err != nil {
				logrus.WithFields(logrus.Fields{}).Error("[Insert] Create contentSchedule error : ", err)
				return errors.New("Lỗi hệ thống vui lòng thử lại")
			}
			contentScheduleD, err := models.ScheduleContents(
				qm.OrderBy("id DESC"),
				qm.Limit(1)).One(ctx, st)
			if err != nil {
				logrus.WithFields(logrus.Fields{}).Error("[ScheduleContents] Find ScheduleContents error : ", err)
				return errors.New("Lỗi hệ thống vui lòng thử lại")
			}
			for _, schedule := range content.SubjectContents {
				scheduleSubject := models.ScheduleSubject{}
				scheduleSubject.Name = schedule.Name
				scheduleSubject.LT = schedule.Lt
				scheduleSubject.TH = schedule.Th
				scheduleSubject.ScheduleSubjectID = contentScheduleD.ID
				err := scheduleSubject.Insert(ctx, st, boil.Infer())
				if err != nil {
					logrus.WithFields(logrus.Fields{}).Error("[Insert] Create scheduleSubject error : ", err)
					return errors.New("Lỗi hệ thống vui lòng thử lại")
				}
			}
		}
	}

	thucHanhs := scheduleResponses.ThucHanh
	for _, thucHanh := range thucHanhs {
		contentSchedule1 := models.ScheduleContent{}
		contentSchedule1.Weekday = thucHanh.WeekDay
		contentSchedule1.Date = thucHanh.Date
		contentSchedule1.SubjectName = null.StringFrom(thucHanh.SubjectName)
		contentSchedule1.HourStudent = null.StringFrom(thucHanh.HourStudent)
		contentSchedule1.KMStudent = null.StringFrom(thucHanh.KmStudent)
		contentSchedule1.HourPerDateVehicle = null.IntFrom(thucHanh.HourPerDateVehicle)
		contentSchedule1.KMDateVehicle = null.IntFrom(thucHanh.KmDateVehicle)
		err := contentSchedule1.Insert(ctx, st, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[Insert] Create contentSchedule error : ", err)
			return errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}
	return nil
}

func generateThucHanhB2(st *sql.DB) ([]ThucHanhResponse, error) {
	thucHanhResponses := []ThucHanhResponse{}
	subjects, err := models.Subjects(
		qm.Where("rank_id = ?", 1),
		qm.And("type = ?", 2),
		qm.OrderBy("id"),
	).All(context.Background(), st)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Subjects] error : ", err)
		return nil, err
	}
	intDays := 0
	for index, subject := range subjects {
		thucHanhResponse := ThucHanhResponse{}
		groups := subject.Group
		if index == 0 {
			startDay := 0
			for i := 0; i < groups; i++ {
				addDate := continueDate.AddDate(0, 0, startDay+1)
				// Ngày học
				thucHanhResponse.Date = utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
				// Nội dung môn học
				thucHanhResponse.SubjectName = subject.Name
				// Thứ trong tuần
				date := addDate.Weekday()
				thucHanhResponse.WeekDay = ConvertDayString(int(date))

				_, ok := holiday[utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)]
				if ok {
					startDay = startDay + 1
					nextDate := continueDate.AddDate(0, 0, startDay)
					_, ok := holiday[utils.TimeIn(nextDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)]
					thucHanhResponse.Date = utils.TimeIn(continueDate.AddDate(0, 0, startDay), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
					date := continueDate.AddDate(0, 0, startDay).Weekday()
					thucHanhResponse.WeekDay = ConvertDayString(int(date))
					if ok {
						next := startDay + 1
						thucHanhResponse.Date = utils.TimeIn(continueDate.AddDate(0, 0, next), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
						date := continueDate.AddDate(0, 0, startDay+1).Weekday()
						thucHanhResponse.WeekDay = ConvertDayString(int(date))
						startDay = next
					}
				}
				// giờ / học viên
				thucHanhResponse.HourStudent = subject.HourStudent.String
				// Km / học viên
				thucHanhResponse.KmStudent = subject.KMStudent.String
				// giờ/ ngày xe
				thucHanhResponse.HourPerDateVehicle = subject.HourDateVehicle.Int
				// Km / Ngày xe
				thucHanhResponse.KmDateVehicle = subject.KMDateVehicle.Int

				thucHanhResponses = append(thucHanhResponses, thucHanhResponse)
				intDays++
				startDay++
			}
		} else {
			ns := groups + intDays
			startDay := intDays
			for i := intDays; i < ns; i++ {
				addDate := continueDate.AddDate(0, 0, startDay)
				if index > 1 {
					addDate = continueDate.AddDate(0, 0, startDay)
				}
				// Ngày học
				thucHanhResponse.Date = utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
				// Nội dung môn học
				thucHanhResponse.SubjectName = subject.Name
				// Thứ trong tuần
				date := addDate.Weekday()
				thucHanhResponse.WeekDay = ConvertDayString(int(date))

				_, ok := holiday[utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)]
				if ok {
					startDay = startDay + 1
					nextDate := continueDate.AddDate(0, 0, startDay)
					_, ok := holiday[utils.TimeIn(nextDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)]
					thucHanhResponse.Date = utils.TimeIn(continueDate.AddDate(0, 0, startDay), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
					date := continueDate.AddDate(0, 0, startDay).Weekday()
					thucHanhResponse.WeekDay = ConvertDayString(int(date))
					if ok {
						next := startDay + 1
						thucHanhResponse.Date = utils.TimeIn(continueDate.AddDate(0, 0, next), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
						date := continueDate.AddDate(0, 0, startDay+1).Weekday()
						thucHanhResponse.WeekDay = ConvertDayString(int(date))
						startDay = next
						intDays++
					}
					intDays++
				}
				// giờ / học viên
				thucHanhResponse.HourStudent = subject.HourStudent.String
				// Km / học viên
				thucHanhResponse.KmStudent = subject.KMStudent.String
				// giờ/ ngày xe
				thucHanhResponse.HourPerDateVehicle = subject.HourDateVehicle.Int
				// Km / Ngày xe
				thucHanhResponse.KmDateVehicle = subject.KMDateVehicle.Int
				thucHanhResponses = append(thucHanhResponses, thucHanhResponse)
				intDays++
				startDay++
			}
		}
	}
	return thucHanhResponses, nil
}

func generateB2(st *sql.DB, course *models.Course) ([]ScheduleResponse, error) {
	scheduleResponses := []ScheduleResponse{}
	from := course.StartDate
	subjects, err := models.Subjects(
		qm.Where("rank_id = ?", 1),
		qm.And("type = ?", 1),
		qm.OrderBy("id"),
	).All(context.Background(), st)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Subjects] error : ", err)
		return nil, err
	}
	intDays := 0
	for index, subject := range subjects {
		scheduleResponse := ScheduleResponse{}
		scheduleResponse.SubjectName = subject.Name
		scheduleResponse.Time = subject.Time
		teacher, err := models.FindTeacher(context.Background(), st, subject.TeacherID.Int)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[FindTeacher] error : ", err)
			return nil, err
		}
		scheduleResponse.Teacher = teacher.Fullname
		groups := subject.Group
		if index == 0 {
			contents := []Content{}
			totalLT := 0
			totalTH := 0
			startDay := 0
			for i := 0; i < groups; i++ {
				content := Content{}
				addDate := from.AddDate(0, 0, startDay)
				content.Date = utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
				date := addDate.Weekday()
				content.WeekDay = ConvertDayString(int(date))

				_, ok := holiday[utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)]
				if ok {
					startDay = startDay + 1
					content.Date = utils.TimeIn(from.AddDate(0, 0, startDay), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
					date := from.AddDate(0, 0, startDay).Weekday()
					content.WeekDay = ConvertDayString(int(date))
					intDays++
				}

				group := i + 1
				childSubject, err := models.ChildSubjects(
					models.ChildSubjectWhere.SubjectID.EQ(subject.ID),
					models.ChildSubjectWhere.Group.EQ(null.StringFrom(strconv.Itoa(group))),
				).All(context.Background(), st)
				if err != nil {
					logrus.WithFields(logrus.Fields{}).Error("[ChildSubjects] error : ", err)
					return nil, err
				}
				if len(childSubject) > 0 {
					subjectContents := []SubjectContent{}
					for _, data := range childSubject {
						subjectContent := SubjectContent{}
						subjectContent.Name = data.Name
						subjectContent.Lt = data.LT.Int
						subjectContent.Th = data.TH.Int
						totalLT += data.LT.Int
						totalTH += data.TH.Int
						subjectContents = append(subjectContents, subjectContent)
					}
					content.SubjectContents = subjectContents
				}
				contents = append(contents, content)
				scheduleResponse.Schedule = contents
				intDays++
				startDay++
			}

			scheduleResponse.TotalLT = totalLT
			scheduleResponse.TotalTH = totalTH
		} else {
			ns := groups + intDays
			groupChild := 1
			contents := []Content{}
			totalLT := 0
			totalTH := 0
			startDay := intDays
			for i := intDays; i < ns; i++ {
				content := Content{}
				addDate := from.AddDate(0, 0, startDay)
				content.Date = utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
				date := addDate.Weekday()
				content.WeekDay = ConvertDayString(int(date))

				_, ok := holiday[utils.TimeIn(addDate, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)]
				if ok {
					startDay = startDay + 1
					content.Date = utils.TimeIn(from.AddDate(0, 0, startDay), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
					date := from.AddDate(0, 0, startDay).Weekday()
					content.WeekDay = ConvertDayString(int(date))
					intDays++
				}
				// convert string to date để lưu lại ngày cuối cùng học lý thuyết
				lastDateLyThuyet, err := time.Parse("02-01-2006", content.Date)
				if err != nil {
					fmt.Println(err)
				}
				continueDate = lastDateLyThuyet

				childSubject, err := models.ChildSubjects(
					models.ChildSubjectWhere.SubjectID.EQ(subject.ID),
					models.ChildSubjectWhere.Group.EQ(null.StringFrom(strconv.Itoa(groupChild))),
				).All(context.Background(), st)
				if err != nil {
					logrus.WithFields(logrus.Fields{}).Error("[ChildSubjects] error : ", err)
					return nil, err
				}
				if len(childSubject) > 0 {
					subjectContents := []SubjectContent{}
					for _, data := range childSubject {
						subjectContent := SubjectContent{}
						subjectContent.Name = data.Name
						subjectContent.Lt = data.LT.Int
						subjectContent.Th = data.TH.Int
						totalLT += data.LT.Int
						totalTH += data.TH.Int
						subjectContents = append(subjectContents, subjectContent)
					}
					content.SubjectContents = subjectContents
				}
				contents = append(contents, content)
				scheduleResponse.Schedule = contents
				intDays++
				groupChild++
				startDay++
			}
			scheduleResponse.TotalLT = totalLT
			scheduleResponse.TotalTH = totalTH
		}
		scheduleResponses = append(scheduleResponses, scheduleResponse)
	}
	logrus.WithFields(logrus.Fields{}).Infoln("[Ngày học lý thuyết cuối cùng]  : ", continueDate)

	return scheduleResponses, nil
}
