package question

import (
	"api-trainning-center/models/admin/result"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (tc StoreQuestion) ShowResult(results []result.Result) (ResponseResult, error) {
	resp := ResponseResult{}
	var countCorrect int
	var countInCorrect int
	resultTests := []ResultTest{}
	for _, result := range results {
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
		resp.NumberOfIncorrect = 35 - countCorrect
		resp.ResultTotal = "KHÔNG ĐẠT - Sai câu điểm liệt"
	} else {
		resp.NumberOfCorrect = countCorrect
		resp.NumberOfIncorrect = 35 - countCorrect
		if resp.NumberOfCorrect >= 32 {
			resp.ResultTotal = "ĐẠT"
		} else {
			resp.ResultTotal = "KHÔNG ĐẠT"
		}
	}
	return resp, nil
}

func checkLiet(db *sql.DB, question, answer string) (bool, error) {
	retrieveLiet, err := retrieveQuestionLiet(db, question, answer)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestionLiet] error : ", err)
		return false, err
	}
	if !retrieveLiet {
		return false, nil
	}
	return true, nil
}

func retrieveQuestionLiet(db *sql.DB, question, answer string) (bool, error) {
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
	row := db.QueryRow(query, question, true)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[checkLiet] Scan error : ", err)
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func retrieveCorrectAnswer(db *sql.DB, question string) (string, error) {
	var result string
	query := `
	select
		q.result
	from
		question q
	where
		q.id =$1
	`
	row := db.QueryRow(query, question)
	err := row.Scan(&result)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveCorrectAnswer] Scan error : ", err)
		return "", err
	}

	return result, nil
}

func checkR(db *sql.DB, question, answer string) (bool, error) {
	retrieveLiet, err := retrieveQuestionR(db, question, answer)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestionR] error : ", err)
		return false, err
	}
	if !retrieveLiet {
		return false, nil
	}
	return true, nil
}

func retrieveQuestionR(db *sql.DB, question, answer string) (bool, error) {
	var count int
	query := `
	select
		count(*)
	from
		question q
	where
		q.id =$1
		and q."result" = $2
	`
	row := db.QueryRow(query, question, answer)
	err := row.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestionR] Scan error : ", err)
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
