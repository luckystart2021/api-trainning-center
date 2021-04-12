package question

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreQuestion) UpdateQuestion(idQuestion int, name, answerA, answerB, answerC, answerD, img, result string, liet bool, typeQ string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := updateQuestionByRequest(st.db, idQuestion, name, answerA, answerB, answerC, answerD, img, result, liet, typeQ); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Cập nhật mật câu hỏi thành công"
	return resp, nil
}

func updateQuestionByRequest(db *sql.DB, idQuestion int, name, answerA, answerB, answerC, answerD, img, result string, liet bool, typeQ string) error {
	if img != "" {
		query :=
			`
	UPDATE question SET 
		name=$1, 
		anwser_correct=$2, 
		paralysis=$3,
		answera=$4,
		answerb=$5,
		answerc=$6,
		answerd=$7,
		img=$8,
		question_type=$9
	WHERE id= $10;
	`
		_, err := db.Exec(query, name, result, liet, answerA, answerB, answerC, answerD, img, typeQ, idQuestion)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateQuestionByRequest] Update DB err  %v", err)
			return errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
	} else {
		query :=
			`
			UPDATE question SET 
				name=$1, 
				anwser_correct=$2, 
				paralysis=$3,
				answera=$4,
				answerb=$5,
				answerc=$6,
				answerd=$7,
				question_type=$8
			WHERE id= $9;
		`
		_, err := db.Exec(query, name, result, liet, answerA, answerB, answerC, answerD, typeQ, idQuestion)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateQuestionByRequest] Update DB err  %v", err)
			return errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
	}

	return nil
}
