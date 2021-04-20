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
	Teacher     string    `json:"teacher"`
	Schedule    []Content `json:"schedule"`
}

type Content struct {
	Date            string           `json:"date"`
	SubjectContents []SubjectContent `json:"subject_contents"`
}

type SubjectContent struct {
	Name string `json:"name"`
	Lt   int    `json:"lt"`
	Th   int    `json:"th"`
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
		teacher, _ := models.FindTeacher(context.Background(), st.db, subject.TeacherID)
		scheduleResponse.Teacher = teacher.Fullname
		groups := subject.Group
		if index == 0 {
			contents := []Content{}
			for i := 0; i < groups; i++ {
				content := Content{}
				content.Date = utils.TimeIn(from.AddDate(0, 0, i), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
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
						subjectContents = append(subjectContents, subjectContent)
					}
					content.SubjectContents = subjectContents
				}
				contents = append(contents, content)
				scheduleResponse.Schedule = contents
				intDays++
			}
		} else {
			ns := groups + intDays
			groupChild := 1
			contents := []Content{}
			for i := intDays; i < ns; i++ {
				content := Content{}
				content.Date = utils.TimeIn(from.AddDate(0, 0, i), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
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
						subjectContents = append(subjectContents, subjectContent)
					}
					content.SubjectContents = subjectContents
				}
				contents = append(contents, content)
				scheduleResponse.Schedule = contents
				intDays++
				groupChild++
			}
		}
		scheduleResponses = append(scheduleResponses, scheduleResponse)
	}
	response.Status = true
	response.Message = "Tạo lịch học thành công"
	return scheduleResponses, nil
}
