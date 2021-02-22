package question

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type Rank struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Time           int    `json:"time"`
	NumberQuestion int    `json:"number_question"`
	PointPass      int    `json:"point_pass"`
}

func (tc StoreQuestion) GetAllTestSuiteByRank(rank string) (TestSuiteResponse, error) {
	testSuite, err := retrieveTestSuite(tc.db, rank)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveTestSuite] error : ", err)
		return TestSuiteResponse{}, err
	}

	rankR, err := retrieveRank(tc.db, rank)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveTimeTest] error : ", err)
		return TestSuiteResponse{}, err
	}
	testSuiteResponse := TestSuiteResponse{}
	testSuiteResponse.Suite = testSuite
	testSuiteResponse.Time = rankR.Time

	return testSuiteResponse, nil
}

func (tc StoreQuestion) GetQuestionsByIdSuite(idSuite int) ([]QuestionResponse, error) {
	questionResponse := []QuestionResponse{}
	questions, err := retrieveQuestion(tc.db, idSuite)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestion] error : ", err)
		return []QuestionResponse{}, err
	}
	for _, data := range questions {
		answers := []Answers{}
		answer1 := Answers{
			Answer: data.AnswerA,
		}
		answers = append(answers, answer1)
		answer2 := Answers{
			Answer: data.AnswerB,
		}
		answers = append(answers, answer2)
		if data.AnswerC != "" || len(data.AnswerC) != 0 {
			answer3 := Answers{
				Answer: data.AnswerC,
			}
			answers = append(answers, answer3)
		}
		if data.AnswerD != "" || len(data.AnswerD) != 0 {
			answer4 := Answers{
				Answer: data.AnswerD,
			}
			answers = append(answers, answer4)
		}

		question := QuestionResponse{
			Id:           data.Id,
			QuestionName: data.Name,
			Img:          data.Img,
			Answers:      answers,
		}
		questionResponse = append(questionResponse, question)
	}

	return questionResponse, nil
}

func retrieveQuestion(db *sql.DB, idSuite int) ([]Questions, error) {
	questions := []Questions{}
	query := `
	select
		q2.id ,
		q2.name ,
		q2.img ,
		q2.answera ,
		q2.answerb ,
		q2.answerc ,
		q2.answerd
	from
		testsuite_question tq
	join question q2 on
		tq.id_question = q2.id
	where
		tq.id_testsuite = $1
	`
	rows, err := db.Query(query, idSuite)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestion] query error  %v", err)
		return questions, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	for rows.Next() {
		var err error
		var id int
		var questionName string
		var answerA, answerB, answerC, answerD, img sql.NullString
		err = rows.Scan(&id, &questionName, &img, &answerA, &answerB, &answerC, &answerD)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestion] Scan error  %v", err)
			return questions, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		question := Questions{
			Id:      id,
			Name:    questionName,
			AnswerA: answerA.String,
			AnswerB: answerB.String,
			AnswerC: answerC.String,
			AnswerD: answerD.String,
		}
		if img.Valid && img.String != "" {
			question.Img = "/files/img/question/" + img.String
		}
		questions = append(questions, question)
	}
	if len(questions) == 0 {
		logrus.WithFields(logrus.Fields{}).Infof("[retrieveQuestion] No Data  %v", err)
		return questions, errors.New("Không có dữ liệu từ hệ thống")
	}
	return questions, nil
}

func retrieveRank(db *sql.DB, rank string) (Rank, error) {
	rankR := Rank{}
	query := `
	SELECT
		id,
		"name",
		"time",
		number_question,
		point_pass
	FROM
		rank_vehicle
	WHERE
		name = $1;
	`
	rows := db.QueryRow(query, rank)
	err := rows.Scan(&rankR.Id, &rankR.Name, &rankR.Time, &rankR.NumberQuestion, &rankR.PointPass)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveRank] No Data  %v", err)
		return rankR, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveRank] Scan error  %v", err)
		return rankR, err
	}
	return rankR, nil
}

func retrieveTestSuite(db *sql.DB, rank string) ([]TestSuite, error) {
	testSuites := []TestSuite{}
	query := `
	SELECT
		t.id ,
		t.name
	FROM
		testsuite t
	JOIN rank_vehicle rv ON
		t.id_rank = rv.id
	WHERE
		rv.name = $1
	`
	rows, err := db.Query(query, rank)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveTestSuite] query error  %v", err)
		return testSuites, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	for rows.Next() {
		var err error
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveTestSuite] Scan error  %v", err)
			return testSuites, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		testSuite := TestSuite{
			Id:   id,
			Name: name,
		}
		testSuites = append(testSuites, testSuite)
	}

	if len(testSuites) == 0 {
		logrus.WithFields(logrus.Fields{}).Infof("[retrieveTestSuite] No Data  %v", err)
		return testSuites, errors.New("Không có dữ liệu từ hệ thống")
	}
	return testSuites, nil
}
