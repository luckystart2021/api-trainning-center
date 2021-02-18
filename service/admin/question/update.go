package question

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreQuestion) UpdateQuestion(idQuestion, codeDe int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error) {
	resp := response.MessageResponse{}

	if err := updateQuestionByRequest(st.db, idQuestion, codeDe, name, answerA, answerB, answerC, answerD, img, result, liet); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Cập nhật mật câu hỏi thành công"
	return resp, nil
}

func updateQuestionByRequest(db *sql.DB, idQuestion, codeDe int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) error {
	query :=
		`
	UPDATE question SET 
		name=$1, 
		result=$2, 
		paralysis=$3,
		id_code_test=$4,
		answera=$5,
		answerb=$6,
		answerc=$7,
		answerd=$8,
		img=$9
	WHERE id= $10;
	`
	_, err := db.Exec(query, name, result, liet, codeDe, answerA, answerB, answerC, answerD, img, idQuestion)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateQuestionByRequest] Update DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
