package question

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreQuestion) CreateQuestion(name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := CreateQuestionByRequest(tc.db, name, answerA, answerB, answerC, answerD, img, result, liet); err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateQuestion]Insert Question DB err  %v", err)
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm câu hỏi thành công"
	return resp, nil
}

func CreateQuestionByRequest(db *sql.DB, name, answerA, answerB, answerC, answerD, img, result string, liet bool) error {
	query := `
	INSERT INTO question
		(name, anwser_correct, paralysis, answera, answerb, answerc, answerd, img)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8);
	`
	_, err := db.Exec(query, name, result, liet, answerA, answerB, answerC, answerD, img)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateInformationByRequest]Insert Question DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
