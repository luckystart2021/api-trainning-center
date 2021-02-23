package suite

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"

	"github.com/sirupsen/logrus"
)

var (
	// Khái niệm
	khainiem = 1
	// Quy tắc giao thông
	quytacgiaothong = 2
	// Tốc độ, khoảng cách
	tocdokhoangcach = 3
	// Chương 2(từ câu số 167 đến câu số 192) Nghiệp vụ vận tải
	nghiepvuvantai = 4
	// Chương 3 (từ câu 193 đến câu 213) Văn hóa giao thông và đạo đức người lái xe
	vanhoagiaothongvadaoducnguoilaixe = 5
	// Chương 4 (từ câu 214 đến câu 269) Kỹ thuật lái xe
	kythuatlaixe = 6
	// Chương 5 (từ câu 270 đến câu 304) Cấu tạo và sửa chữa
	cautaovasuachua = 7
	// Chương 6 (từ số 305 đến câu số 486). Hệ thống biển báo hiệu đường bộ
	hethongbienbaohieuduongbo = 8
	// Chương 7 (từ số 487 đến câu số 600). Các thế sa hình và kỹ năng xử lý tình huống giao thông
	cacthexahinhvakynangxulytinhhuonggiaothong = 9
)

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func randomQuestion(input []QuestionResp) []int {
	fmt.Println("len", len(input))
	ir := IntRange{1, len(input)}
	var idQuestRans []int
	var id int
	var n = 100
	for j := 0; j < n; j++ {
		idRan := rand.Intn(1000)
		r := rand.New(rand.NewSource(int64(idRan)))
		id = ir.NextRandom(r)
		for _, data1 := range input {
			if data1.Id == id {
				idQuestRans = append(idQuestRans, data1.IdQuestion)
			}
		}
		if checkOverlap(idQuestRans) {
			idQuestRans = idQuestRans[:len(idQuestRans)-1]
			continue
		}
		if len(idQuestRans) == 10 {
			break
		}
	}
	fmt.Println("ket qua ran", idQuestRans)
	return idQuestRans
}

func checkOverlap(as []int) bool {
	for i := 0; i < len(as)-1; i++ {
		for j := i + 1; j < len(as); j++ {
			if as[i] == as[j] {
				return true
			}
		}
	}
	return false
}

func (tc StoreSuiteTest) GenarateSuiteTest(number, rank string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	questions, err := retrieveQuestion(tc.db, khainiem)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestion] error : ", err)
		return resp, err
	}
	fmt.Println("questions", questions)
	randomQuestion(questions)
	return resp, nil
}

type QuestionResp struct {
	Id         int
	IdQuestion int
}

func retrieveQuestion(db *sql.DB, questionType int) ([]QuestionResp, error) {
	res := []QuestionResp{}
	query := `
	select 
		q.id
	from 
		question q 
	where 
		question_type = $1
	`
	rows, err := db.Query(query, questionType)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestion] query error  %v", err)
		return res, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	rs := res
	index := 0
	for rows.Next() {
		index++
		var r QuestionResp
		err = rows.Scan(&r.IdQuestion)
		r.Id = index
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestion] Scan error  %v", err)
			return res, errors.New("Lỗi hệ thống vui lòng thử lại")
		}

		rs = append(rs, r)
	}

	return rs, nil
}

func insertSuiteTestAndQuestion(db *sql.DB, idSuiteTest, idQuestion int) int64 {
	query := `
	INSERT INTO testsuite_question
		(id_testsuite, id_question)
	VALUES($1, $2);
	`
	res, err := db.Exec(query, idSuiteTest, idQuestion)
	rowsAffected, _ := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[insertSuiteTestAndQuestion]Insert TestSuite_Question DB err  %v", err)
		return 0
	}
	return rowsAffected
}
