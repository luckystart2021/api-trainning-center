package question

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreQuestion) ShowQuestionsExam() ([]ResponseQuestionExam, error) {
	questionsExam, err := retrieveShowQuestionsExam(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowQuestionsExam] error : ", err)
		return []ResponseQuestionExam{}, err
	}

	return questionsExam, nil
}

func retrieveShowQuestionsExam(db *sql.DB) ([]ResponseQuestionExam, error) {
	questionsExam := []ResponseQuestionExam{}
	query := `
	SELECT id, name
	FROM testsuite
	order by id;
	`
	rows, err := db.Query(query)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveShowQuestionsExam] No Data  %v", err)
		return questionsExam, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveShowQuestionsExam] query error  %v", err)
		return questionsExam, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		var err error
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveShowQuestionsExam] Scan error  %v", err)
			return questionsExam, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		questionExam := ResponseQuestionExam{
			Id:   id,
			Name: name,
		}
		questionsExam = append(questionsExam, questionExam)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveShowQuestionsExam] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(questionsExam) == 0 {
		logrus.WithFields(logrus.Fields{}).Infof("[retrieveShowQuestionsExam] No Data  %v", err)
		return questionsExam, errors.New("Không có dữ liệu từ hệ thống")
	}
	return questionsExam, nil
}

func (tc StoreQuestion) ShowQuestions(code string) ([]Question, error) {
	question, err := retrieveQuestions(tc.db, code)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowQuestions] error : ", err)
		return []Question{}, err
	}

	return question, nil
}

func (tc StoreQuestion) ShowQuestionsSystem() ([]QuestionSystem, error) {
	question, err := retrieveQuestionsSystem(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowQuestionsSystem] error : ", err)
		return []QuestionSystem{}, err
	}

	return question, nil
}

func (tc StoreQuestion) ShowQuestionSystem(idQuestion string) (QuestionSystem, error) {
	question, err := retrieveQuestionSystem(tc.db, idQuestion)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowQuestionSystem] error : ", err)
		return QuestionSystem{}, err
	}

	return question, nil
}

func retrieveQuestionSystem(db *sql.DB, idQuestion string) (QuestionSystem, error) {
	question := QuestionSystem{}
	query := `
	select
		q.id,
		q.name,
		q.anwser_correct,
		q.paralysis,
		q.answera ,
		q.answerb ,
		q.answerc ,
		q.answerd,
		q.img
	from
		question q
	where
		q.id = $1
	`
	rows := db.QueryRow(query, idQuestion)
	var id int64
	var paralysis bool
	var questionName, result string
	var answerA, answerB, answerC, answerD, img sql.NullString
	err := rows.Scan(&id, &questionName, &result, &paralysis, &answerA, &answerB, &answerC, &answerD, &img)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionSystem] No Data  %v", err)
		return question, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionSystem] Scan error  %v", err)
		return question, err
	}
	questionSystem := QuestionSystem{
		Id:            id,
		Name:          questionName,
		AnswerA:       answerA.String,
		AnswerB:       answerB.String,
		AnswerC:       answerC.String,
		AnswerD:       answerD.String,
		AnwserCorrect: result,
		Liet:          paralysis,
	}
	if img.Valid && img.String != "" {
		questionSystem.Img = "/files/img/question/" + img.String
	}
	return questionSystem, nil
}

func retrieveQuestionsSystem(db *sql.DB) ([]QuestionSystem, error) {
	questions := []QuestionSystem{}
	query := `
	select
		q.id,
		q.name,
		q.anwser_correct,
		q.paralysis,
		q.answera ,
		q.answerb ,
		q.answerc ,
		q.answerd,
		q.img
	from
		question q
	order by id;
	`
	rows, err := db.Query(query)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionsSystem] No Data  %v", err)
		return questions, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionsSystem] query error  %v", err)
		return questions, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		var err error
		var id int64
		var paralysis bool
		var questionName, result string
		var answerA, answerB, answerC, answerD, img sql.NullString
		err = rows.Scan(&id, &questionName, &result, &paralysis, &answerA, &answerB, &answerC, &answerD, &img)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionsSystem] Scan error  %v", err)
			return questions, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		questionSystem := QuestionSystem{
			Id:            id,
			Name:          questionName,
			AnswerA:       answerA.String,
			AnswerB:       answerB.String,
			AnswerC:       answerC.String,
			AnswerD:       answerD.String,
			AnwserCorrect: result,
			Liet:          paralysis,
		}
		if img.Valid && img.String != "" {
			questionSystem.Img = "/files/img/question/" + img.String
		}
		questions = append(questions, questionSystem)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionsSystem] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(questions) == 0 {
		logrus.WithFields(logrus.Fields{}).Infof("[retrieveQuestionsSystem] No Data  %v", err)
		return questions, errors.New("Không có dữ liệu từ hệ thống")
	}
	return questions, nil
}

func retrieveQuestions(db *sql.DB, code string) ([]Question, error) {
	questions := []Question{}
	query := `
	select
		q.id,
		q.name,
		q.answera ,
		q.answerb ,
		q.answerc ,
		q.answerd,
		q.img
	from
		question q
	inner join testsuite t on
		t.id = q.id_code_test
	where
		q.id_code_test = $1
	order by id;
	`
	rows, err := db.Query(query, code)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestions] No Data  %v", err)
		return questions, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestions] query error  %v", err)
		return questions, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		var err error
		var questionName, id string
		var answerA, answerB, answerC, answerD, img sql.NullString
		err = rows.Scan(&id, &questionName, &answerA, &answerB, &answerC, &answerD, &img)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestions] Scan error  %v", err)
			return questions, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		answers := []Answers{}
		answer1 := Answers{
			Answer: answerA.String,
		}
		answer2 := Answers{
			Answer: answerB.String,
		}
		answer3 := Answers{
			Answer: answerC.String,
		}
		answer4 := Answers{
			Answer: answerD.String,
		}

		if !answerC.Valid && !answerD.Valid {
			answers = append(answers, answer1, answer2)
		} else if !answerD.Valid {
			answers = append(answers, answer1, answer2, answer3)
		} else {
			answers = append(answers, answer1, answer2, answer3, answer4)
		}

		question := Question{
			Id:           id,
			QuestionName: questionName,
			Answers:      answers,
		}
		if img.Valid && img.String != "" {
			question.Img = "/files/img/question/" + img.String
		}
		questions = append(questions, question)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestions] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(questions) == 0 {
		logrus.WithFields(logrus.Fields{}).Infof("[retrieveQuestions] No Data  %v", err)
		return questions, errors.New("Không có dữ liệu từ hệ thống")
	}
	return questions, nil
}
