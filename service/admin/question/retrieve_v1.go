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
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveTestSuite] No Data  %v", err)
		return testSuites, errors.New("Không có dữ liệu từ hệ thống")
	}
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
