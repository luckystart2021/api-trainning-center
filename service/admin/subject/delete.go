package subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"
	"log"

	"github.com/sirupsen/logrus"
)

func (st StoreSubject) DeleteSubject(idSubject int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	childSubject, err := models.ChildSubjects(
		models.ChildSubjectWhere.SubjectID.EQ(idSubject),
	).All(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindChildSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}
	_, err = childSubject.DeleteAll(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteAllChildSubject] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	subject, err := models.FindSubject(ctx, tx, idSubject)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindSubject] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}
	rowsAff, err := subject.Delete(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteSubject] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
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
