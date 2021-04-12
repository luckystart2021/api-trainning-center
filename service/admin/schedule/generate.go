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
)

type ScheduleResponse struct {
	Date           string   `json:"date"`
	SubjectName    string   `json:"subject_name"`
	SubjectContent []string `json:"subject_content"`
}

func (st StoreSchedule) GenerateSchedule(courseId int) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	scheduleResponses := []ScheduleResponse{}
	course, _ := models.FindCourse(context.Background(), st.db, courseId)

	to := course.EndDate
	from := course.StartDate
	// fmt.Println("SSS", from.AddDate(0, 3, 0))
	days := to.Sub(from) / (24 * time.Hour)
	fmt.Println("days", int(days))
	noofdays := int(days)

	for i := 0; i <= noofdays; i++ {
		scheduleResponse := ScheduleResponse{}
		scheduleResponse.Date = utils.TimeIn(from.AddDate(0, 0, i), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY)
		// fmt.Println(utils.TimeIn(from.AddDate(0, 0, i), utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYY))
		scheduleResponses = append(scheduleResponses, scheduleResponse)
	}

	for i := range scheduleResponses {
		// subject, _ := models.Subjects().All(context.Background(), st.db)
		// for _, data := range subject {
		index := strconv.Itoa(i + 1)
		childSubject, _ := models.ChildSubjects(
			models.ChildSubjectWhere.Group.EQ(null.StringFrom(index)),
		).All(context.Background(), st.db)
		for _, dataS := range childSubject {
			scheduleResponses[i].SubjectContent = append(scheduleResponses[i].SubjectContent, dataS.Name)
		}
	}
	fmt.Println("scheduleResponses", scheduleResponses)
	response.Status = true
	response.Message = "Tạo lịch học thành công"
	return response, nil
}
