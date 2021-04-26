package schedule

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

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
	// khai báo một mảng các string
	// toán tử ... để đếm số phần tử
	// số phần tử của mảng là (7)
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

var holiday = map[string]string{
	// "26-03-2020": "le",
	// "30-03-2020": "le",
	// "19-04-2020": "le",
}

func (st StoreSchedule) GenerateSchedule(courseId int) ([]ScheduleResponse, error) {
	response := response.MessageResponse{}
	scheduleResponses := []ScheduleResponse{}
	course, _ := models.FindCourse(context.Background(), st.db, courseId)

	to := course.EndDate
	from := course.StartDate
	// fmt.Println("SSS", from.AddDate(0, 3, 0))
	days := to.Sub(from) / (24 * time.Hour)
	fmt.Println("days", int(days))
	// noofdays := int(days)
	subjects, _ := models.Subjects(
		qm.OrderBy("id"),
	).All(context.Background(), st.db)
	intDays := 0
	for index, subject := range subjects {
		scheduleResponse := ScheduleResponse{}
		scheduleResponse.SubjectName = subject.Name
		scheduleResponse.Time = subject.Time
		teacher, _ := models.FindTeacher(context.Background(), st.db, subject.TeacherID)
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
				childSubject, _ := models.ChildSubjects(
					models.ChildSubjectWhere.SubjectID.EQ(subject.ID),
					models.ChildSubjectWhere.Group.EQ(null.StringFrom(strconv.Itoa(group))),
				).All(context.Background(), st.db)
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

				childSubject, _ := models.ChildSubjects(
					models.ChildSubjectWhere.SubjectID.EQ(subject.ID),
					models.ChildSubjectWhere.Group.EQ(null.StringFrom(strconv.Itoa(groupChild))),
				).All(context.Background(), st.db)
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
	response.Status = true
	response.Message = "Tạo lịch học thành công"
	return scheduleResponses, nil
}
