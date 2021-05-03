package subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreSubject) UpdateSubject(subjectID int, req models.Subject, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	subjectFromDB, err := models.FindSubject(ctx, st.db, subjectID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	if req.Type.Int == 2 {
		subjectFromDB.Name = req.Name
		subjectFromDB.UpdatedBy = userName
		subjectFromDB.Group = req.Group
		subjectFromDB.RankID = req.RankID
		subjectFromDB.Type = req.Type
		subjectFromDB.HourStudent = req.HourStudent
		subjectFromDB.KMStudent = req.KMStudent
		subjectFromDB.HourDateVehicle = req.HourDateVehicle
		subjectFromDB.KMDateVehicle = req.KMDateVehicle
	}
	if req.Type.Int == 1 {
		subjectFromDB.Name = req.Name
		subjectFromDB.Time = req.Time
		subjectFromDB.TeacherID = req.TeacherID
		subjectFromDB.UpdatedBy = userName
		subjectFromDB.Group = req.Group
		subjectFromDB.RankID = req.RankID
		subjectFromDB.Type = req.Type
	}

	rowsAff, err := subjectFromDB.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateSubject] Update Subject error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin môn học thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin môn học không thành công"
	}
	return resp, nil
}
