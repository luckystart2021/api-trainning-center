package question

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreQuestion) UpdateQuestion(idQuestion int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error) {
	resp := response.MessageResponse{}

	if err := updateQuestionByRequest(st.db, idQuestion, name, answerA, answerB, answerC, answerD, img, result, liet); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Cập nhật mật câu hỏi thành công"
	return resp, nil
}

func updateQuestionByRequest(db *sql.DB, idQuestion int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) error {
	if img != "" {
		query :=
			`
	UPDATE question SET 
		name=$1, 
		result=$2, 
		paralysis=$3,
		answera=$4,
		answerb=$5,
		answerc=$6,
		answerd=$7,
		img=$9
	WHERE id= $9;
	`
		_, err := db.Exec(query, name, result, liet, answerA, answerB, answerC, answerD, img, idQuestion)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateQuestionByRequest] Update DB err  %v", err)
			return errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
	} else {
		query :=
			`
			UPDATE question SET 
				name=$1, 
				result=$2, 
				paralysis=$3,
				answera=$4,
				answerb=$5,
				answerc=$6,
				answerd=$7
			WHERE id= $8;
		`
		_, err := db.Exec(query, name, result, liet, answerA, answerB, answerC, answerD, idQuestion)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateQuestionByRequest] Update DB err  %v", err)
			return errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
	}

	return nil
}
