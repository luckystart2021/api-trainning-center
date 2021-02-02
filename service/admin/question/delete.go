package question

import (
	"api-trainning-center/service/response"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (tc StoreQuestion) DeleteQuestion(idQuestion string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := deleteQuestionByID(tc.db, idQuestion)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[DeleteQuestion] Error %v", err)
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa câu hỏi thành công"
	} else {
		resp.Status = false
		resp.Message = "Không có câu hỏi được xóa"
	}
	return resp, nil
}

// delete user in the DB
func deleteQuestionByID(db *sql.DB, idQuestion string) (int64, error) {
	query := `
	DELETE 
	FROM 
		question
	WHERE 
		id=$1;
	`
	res, err := db.Exec(query, idQuestion)
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteQudeleteQuestionByIDestion] Error while checking the affected rows. %v", err)
		return 0, err
	}
	logrus.WithFields(logrus.Fields{}).Infof("[deleteQuestionByID] Total rows/record affected %v", rowsAffected)
	return rowsAffected, nil
}
