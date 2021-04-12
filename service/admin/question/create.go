package question

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (tc StoreQuestion) CreateQuestion(name, answerA, answerB, answerC, answerD, img, result string, liet bool, typeQ string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()
	questionId, err := models.Questions(
		qm.OrderBy("id DESC"),
		qm.Limit(1),
	).One(ctx, tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[models.Questions]Get id Question DB err  %v and ID %s", err, questionId.ID)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	if err := CreateQuestionByRequest(tc.db, questionId.ID, name, answerA, answerB, answerC, answerD, img, result, liet, typeQ); err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateQuestion]Insert Question DB err  %v", err)
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm câu hỏi thành công"
	return resp, nil
}

func CreateQuestionByRequest(db *sql.DB, id int, name, answerA, answerB, answerC, answerD, img, result string, liet bool, typeQ string) error {
	query := `
	INSERT INTO question
		(id ,name, anwser_correct, paralysis, answera, answerb, answerc, answerd, img, question_type)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`
	_, err := db.Exec(query, id+1, name, result, liet, answerA, answerB, answerC, answerD, img, typeQ)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateInformationByRequest]Insert Question DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
