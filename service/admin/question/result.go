package question

import (
	"api-trainning-center/models/admin/result"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreQuestion) ShowResult(results result.Result) (ResponseResult, error) {
	resp := ResponseResult{}
	numberAndPointPass, err := retrieveNumberAndPointPass(tc.db, results.IdSuite)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveNumberAndPointPass] error : ", err)
		return ResponseResult{}, err
	}
	if len(results.Answers) > numberAndPointPass.NumberQuestion {
		logrus.WithFields(logrus.Fields{}).Error("[ShowResult] error : ", err)
		return ResponseResult{}, errors.New("Câu trả lời vượt quá số câu qui định")
	}
	var countCorrect int
	var countInCorrect int
	resultTests := []ResultTest{}
	for _, result := range results.Answers {
		checkR, err := checkR(tc.db, result.IdQuestion, result.IdAnswer)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[checkR] error : ", err)
			return ResponseResult{}, err
		}
		if checkR {
			countCorrect++
		} else {
			resultLiet, err := checkLiet(tc.db, result.IdQuestion, result.IdAnswer)
			if err != nil {
				logrus.WithFields(logrus.Fields{}).Error("[checkLiet] error : ", err)
				return ResponseResult{}, err
			}
			if resultLiet {
				countInCorrect++
			}
		}
		resultFromDB, err := retrieveCorrectAnswer(tc.db, result.IdQuestion)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[retrieveCorrectAnswer] error : ", err)
			return ResponseResult{}, err
		}
		resultTest := ResultTest{
			IdQuestion:    result.IdQuestion,
			IdAnswer:      result.IdAnswer,
			CorrectAnswer: resultFromDB,
		}
		resultTests = append(resultTests, resultTest)
	}
	resp.ResultTests = resultTests
	if countInCorrect >= 1 {
		resp.NumberOfCorrect = countCorrect
		resp.NumberOfIncorrect = numberAndPointPass.NumberQuestion - countCorrect
		resp.ResultTotal = "KHÔNG ĐẠT - Sai câu điểm liệt"
	} else {
		resp.NumberOfCorrect = countCorrect
		resp.NumberOfIncorrect = numberAndPointPass.NumberQuestion - countCorrect
		if resp.NumberOfCorrect >= numberAndPointPass.PointPass {
			resp.ResultTotal = "ĐẠT"
		} else {
			resp.ResultTotal = "KHÔNG ĐẠT"
		}
	}
	return resp, nil
}

type NumberAndPointPassResponse struct {
	NumberQuestion int
	PointPass      int
}

func retrieveNumberAndPointPass(db *sql.DB, idSuite int) (NumberAndPointPassResponse, error) {
	resp := NumberAndPointPassResponse{}
	query := `
	SELECT
		rv.number_question ,
		rv.point_pass 
	FROM
		testsuite t
	INNER JOIN rank_vehicle rv ON
		rv.id = t.id_rank
	WHERE t.id = $1
	`
	row := db.QueryRow(query, idSuite)
	err := row.Scan(&resp.NumberQuestion, &resp.PointPass)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveNumberAndPointPass] No Data  %v", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveNumberAndPointPass] Scan error : ", err)
		return resp, err
	}
	return resp, nil
}

func checkLiet(db *sql.DB, idQuestion int, answer string) (bool, error) {
	retrieveLiet, err := retrieveQuestionLiet(db, idQuestion, answer)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestionLiet] error : ", err)
		return false, err
	}
	if !retrieveLiet {
		return false, nil
	}
	return true, nil
}

func retrieveQuestionLiet(db *sql.DB, idQuestion int, answer string) (bool, error) {
	var count int
	query := `
	select
		count(*)
	from
		question q
	where
		q.id =$1
		and q.paralysis = $2
	`
	row := db.QueryRow(query, idQuestion, true)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[checkLiet] Scan error : ", err)
		return false, errors.New("Không có dữ liệu từ hệ thống")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func retrieveCorrectAnswer(db *sql.DB, idQuestion int) (string, error) {
	var result string
	query := `
	select
		q.anwser_correct
	from
		question q
	where
		q.id =$1
	`
	row := db.QueryRow(query, idQuestion)
	err := row.Scan(&result)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveCorrectAnswer] Scan error : ", err)
		return "", errors.New("Không có dữ liệu từ hệ thống")
	}

	return result, nil
}

func checkR(db *sql.DB, idQuestion int, answer string) (bool, error) {
	retrieveLiet, err := retrieveQuestionR(db, idQuestion, answer)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestionR] error : ", err)
		return false, err
	}
	if !retrieveLiet {
		return false, nil
	}
	return true, nil
}

func retrieveQuestionR(db *sql.DB, idQuestion int, answer string) (bool, error) {
	var count int
	query := `
	select
		count(*)
	from
		question q
	where
		q.id =$1
		and q.anwser_correct = $2
	`
	row := db.QueryRow(query, idQuestion, answer)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestionR] Scan error : ", err)
		return false, errors.New("Không có dữ liệu từ hệ thống")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
