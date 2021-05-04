package child_subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"
	"log"

	"github.com/sirupsen/logrus"
)

func (st StoreChildSubject) DeleteChildSubject(childSubjectID int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	childSubject, err := models.FindChildSubject(ctx, tx, childSubjectID)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindChildSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	rowsAff, err := childSubject.Delete(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteChildSubject] error : ", err)
		return resp, errors.New("Hệ thống lỗi vui lòng thử lại")
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Xóa thành công"
	} else {
		resp.Status = false
		resp.Message = "Xóa không thành công"
	}
	return resp, nil
}
