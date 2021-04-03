package teacher

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreTeacher) CreateTeacher(req models.Teacher, userName string) (response.MessageResponse, error) {
	ctx := context.Background()
	resp := response.MessageResponse{}
	req.CreatedBy = userName
	req.UpdatedBy = userName
	err := req.Insert(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CreateTeacher] Create Teacher error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	resp.Status = true
	resp.Message = "Thêm giáo viên thành công"
	return resp, nil
}
