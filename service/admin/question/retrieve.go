package question

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (tc StoreQuestion) ShowQuestions(code string) ([]Question, error) {
	question, err := retrieveQuestions(tc.db, code)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowQuestions] error : ", err)
		return []Question{}, err
	}

	return question, nil
}

func (tc StoreQuestion) ShowQuestionsSystem(code string) ([]QuestionSystem, error) {
	question, err := retrieveQuestionsSystem(tc.db, code)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowQuestionsSystem] error : ", err)
		return []QuestionSystem{}, err
	}

	return question, nil
}

func retrieveQuestionsSystem(db *sql.DB, code string) ([]QuestionSystem, error) {
	questions := []QuestionSystem{}
	query := `
	select
		q.id,
		q.name,
		q.result,
		q.id_code_test,
		q.paralysis,
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
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionsSystem] query error  %v", err)
		return questions, err
	}

	for rows.Next() {
		var err error
		var id, codeDe int64
		var paralysis bool
		var questionName, result string
		var answerA, answerB, answerC, answerD, img sql.NullString
		err = rows.Scan(&id, &questionName, &result, &codeDe, &paralysis, &answerA, &answerB, &answerC, &answerD, &img)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionsSystem] Scan error  %v", err)
			return questions, err
		}
		questionSystem := QuestionSystem{
			Id:      id,
			CodeDe:  codeDe,
			Name:    questionName,
			AnswerA: answerA.String,
			AnswerB: answerB.String,
			AnswerC: answerC.String,
			AnswerD: answerD.String,
			Img:     img.String,
			Result:  result,
			Liet:    paralysis,
		}
		questions = append(questions, questionSystem)
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
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestions] query error  %v", err)
		return questions, err
	}

	for rows.Next() {
		var err error
		var questionName, id string
		var answerA, answerB, answerC, answerD, img sql.NullString
		err = rows.Scan(&id, &questionName, &answerA, &answerB, &answerC, &answerD, &img)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestions] Scan error  %v", err)
			return questions, err
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
			Img:          "/files/img/question/" + img.String,
			Answers:      answers,
		}
		questions = append(questions, question)
	}
	return questions, nil
}
