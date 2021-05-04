package child_subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreChildSubject) UpdateChildSubject(childSubjectID int, req models.ChildSubject, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	childSubjectFromDB, err := models.FindChildSubject(ctx, st.db, childSubjectID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindChildSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	subject, err := models.Subjects(
		models.SubjectWhere.ID.EQ(req.SubjectID),
	).Exists(context.Background(), st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}
	if !subject {
		logrus.WithFields(logrus.Fields{}).Error("[FindSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu môn học từ hệ thống")
	}

	childSubjectFromDB.Name = req.Name
	childSubjectFromDB.LT = req.LT
	childSubjectFromDB.TH = req.TH
	childSubjectFromDB.Group = req.Group
	childSubjectFromDB.SubjectID = req.SubjectID
	childSubjectFromDB.UpdatedBy = userName

	rowsAff, err := childSubjectFromDB.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateChildSubject] Update Child Subject error : ", err)
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
