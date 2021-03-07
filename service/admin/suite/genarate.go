package suite

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"

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
	// các câu liệt
	liet = 10
)

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func randomQuestion(input []QuestionResp, loaicauhoi int, rank string) []int {
	ir := IntRange{1, len(input)}
	var idQuestRans []int
	var id int
	var n int
	if loaicauhoi == khainiem {
		n = 1
	}
	if loaicauhoi == quytacgiaothong {
		n = 7
	}
	if loaicauhoi == tocdokhoangcach {
		n = 1
	}
	if loaicauhoi == nghiepvuvantai {
		n = 1
	}
	if loaicauhoi == vanhoagiaothongvadaoducnguoilaixe {
		n = 1
	}
	if loaicauhoi == kythuatlaixe {
		n = 2
	}
	if loaicauhoi == cautaovasuachua {
		n = 1
	}
	if loaicauhoi == hethongbienbaohieuduongbo && rank == "B2" {
		n = 10
	}
	if loaicauhoi == hethongbienbaohieuduongbo && rank == "C" {
		n = 14
	}
	if loaicauhoi == hethongbienbaohieuduongbo && rank == "DEF" {
		n = 16
	}

	if loaicauhoi == cacthexahinhvakynangxulytinhhuonggiaothong && rank == "B2" {
		n = 10
	}
	if loaicauhoi == cacthexahinhvakynangxulytinhhuonggiaothong && rank == "C" {
		n = 11
	}
	if loaicauhoi == cacthexahinhvakynangxulytinhhuonggiaothong && rank == "DEF" {
		n = 14
	}
	if loaicauhoi == liet {
		n = 1
	}

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
			j--
			continue
		}
	}

	if len(idQuestRans) != n {
		logrus.WithFields(logrus.Fields{}).Error("Chưa đủ bộ đề")
	}
	fmt.Println("toi day r bne 2", idQuestRans)

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

var khongliet = false

func (tc StoreSuiteTest) GenarateSuiteTest(number, rank string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	idR, _ := strconv.Atoi(number)
	insertSuiteTest(tc.db, idR, rank)
	boDe := retrieveBoDe(tc.db, rank)
	for _, bd := range boDe {
		// for index := 0; index < idR; index++ {
		questionsKhaiNiem, err := retrieveQuestion(tc.db, khainiem, false)
		questionsQuytacgiathong, err := retrieveQuestion(tc.db, quytacgiaothong, false)
		questionsTocdokhoangcach, err := retrieveQuestion(tc.db, tocdokhoangcach, false)
		questionsNghiepvuvantai, err := retrieveQuestion(tc.db, nghiepvuvantai, false)
		questionsVanhoavadaoduc, err := retrieveQuestion(tc.db, vanhoagiaothongvadaoducnguoilaixe, false)
		questionsKythuatlaixe, err := retrieveQuestion(tc.db, kythuatlaixe, false)
		questionsCautaovasuachua, err := retrieveQuestion(tc.db, cautaovasuachua, false)
		questionsHethongbienbaohieuduongbo, err := retrieveQuestion(tc.db, hethongbienbaohieuduongbo, false)
		questionsCacthexahinhvakynangxulytinhhuonggiaothong, err := retrieveQuestion(tc.db, cacthexahinhvakynangxulytinhhuonggiaothong, false)
		questionsLiet, err := retrieveQuestionLiet(tc.db)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[retrieveQuestion] error : ", err)
			return resp, err
		}
		total := []int{}
		totalKhaiNiem := randomQuestion(questionsKhaiNiem, khainiem, rank)
		for _, data := range totalKhaiNiem {
			total = append(total, data)
		}
		totalQuytacgiathong := randomQuestion(questionsQuytacgiathong, quytacgiaothong, rank)
		for _, data := range totalQuytacgiathong {
			total = append(total, data)
		}
		totalTocdokhoangcach := randomQuestion(questionsTocdokhoangcach, tocdokhoangcach, rank)
		for _, data := range totalTocdokhoangcach {
			total = append(total, data)
		}

		totalNghiepvuvantai := randomQuestion(questionsNghiepvuvantai, nghiepvuvantai, rank)
		for _, data := range totalNghiepvuvantai {
			total = append(total, data)
		}

		totalVanhoavadaoduc := randomQuestion(questionsVanhoavadaoduc, vanhoagiaothongvadaoducnguoilaixe, rank)
		for _, data := range totalVanhoavadaoduc {
			total = append(total, data)
		}

		totalKythuatlaixe := randomQuestion(questionsKythuatlaixe, kythuatlaixe, rank)
		for _, data := range totalKythuatlaixe {
			total = append(total, data)
		}

		totalCautaovasuachua := randomQuestion(questionsCautaovasuachua, cautaovasuachua, rank)
		for _, data := range totalCautaovasuachua {
			total = append(total, data)
		}

		totalHethongbienbaohieuduongbo := randomQuestion(questionsHethongbienbaohieuduongbo, hethongbienbaohieuduongbo, rank)
		for _, data := range totalHethongbienbaohieuduongbo {
			total = append(total, data)
		}

		totalCacthexahinhvakynangxulytinhhuonggiaothong := randomQuestion(questionsCacthexahinhvakynangxulytinhhuonggiaothong, cacthexahinhvakynangxulytinhhuonggiaothong, rank)
		for _, data := range totalCacthexahinhvakynangxulytinhhuonggiaothong {
			total = append(total, data)
		}

		totalLiet := randomQuestion(questionsLiet, liet, rank)
		for _, data := range totalLiet {
			total = append(total, data)
		}

		if rank == "B2" && len(total) != 35 {
			fmt.Println("B2 cần 35 câu hỏi")
			resp.Message = "thất bại"
			return resp, nil
		}
		if rank == "C" && len(total) != 40 {
			fmt.Println("C cần 40 câu hỏi")
			resp.Message = "thất bại"
			return resp, nil
		}
		if rank == "DEF" && len(total) != 45 {
			fmt.Println("DEFcần 45 câu hỏi")
			resp.Message = "thất bại"
			return resp, nil
		}

		after := []int{}

		for j := 0; j < len(total); j++ {
			randomIndex := rand.Intn(len(total))
			dataR := total[randomIndex]
			after = append(after, dataR)

			if checkOverlap(after) {
				after = after[:len(after)-1]
				j--
				continue
			}
		}

		for _, data := range after {
			insertSuiteTestAndQuestion(tc.db, bd, data)
		}

		resp.Status = true
		resp.Message = "thành công"
		logrus.WithFields(logrus.Fields{}).Info("[GenarateSuiteTest] tổng số câu sau khi random", len(after))
		logrus.WithFields(logrus.Fields{}).Info("[GenarateSuiteTest] ID", after)
		// }
	}
	return resp, nil
}

type QuestionResp struct {
	Id         int
	IdQuestion int
}

func retrieveQuestion(db *sql.DB, questionType int, liet bool) ([]QuestionResp, error) {
	res := []QuestionResp{}
	query := `
	select 
		q.id
	from 
		question q 
	where 
		question_type = $1
		and paralysis = $2
	`
	rows, err := db.Query(query, questionType, liet)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestion] query error  %v", err)
		return res, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	rs := res
	index := 0
	defer rows.Close()
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

func retrieveQuestionLiet(db *sql.DB) ([]QuestionResp, error) {
	res := []QuestionResp{}
	query := `
	select 
		q.id
	from 
		question q 
	where 
		 paralysis = $1
	`
	rows, err := db.Query(query, true)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionLiet] query error  %v", err)
		return res, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	rs := res
	index := 0
	defer rows.Close()
	for rows.Next() {
		index++
		var r QuestionResp
		err = rows.Scan(&r.IdQuestion)
		r.Id = index
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveQuestionLiet] Scan error  %v", err)
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

func insertSuiteTest(db *sql.DB, number int, rank string) int64 {
	var rowsAffected int64
	var idRank int
	if rank == "B2" {
		idRank = 1
	} else if rank == "C" {
		idRank = 2
	} else {
		idRank = 3
	}
	for j := 0; j < number; j++ {

		idDe := strconv.Itoa(j + 1)
		maDe := "Đề " + idDe
		query := `
		INSERT INTO testsuite(name, id_rank)
		VALUES($1, $2);
		`
		res, err := db.Exec(query, maDe, idRank)
		rowsAffected, _ = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[insertSuiteTest]Insert TestSuite DB err  %v", err)
			return 0
		}
	}
	return rowsAffected
}

func retrieveBoDe(db *sql.DB, rank string) []int {
	var resp []int
	var idRank int
	if rank == "B2" {
		idRank = 1
	} else if rank == "C" {
		idRank = 2
	} else {
		idRank = 3
	}
	query := `
	select 
		id
	from 
		testsuite 
	where 
		id_rank = $1
	`
	rows, err := db.Query(query, idRank)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveBoDe] query error  %v", err)
		return resp
	}
	defer rows.Close()
	for rows.Next() {
		var r int
		err = rows.Scan(&r)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveBoDe] Scan error  %v", err)
			return resp
		}

		resp = append(resp, r)
	}
	return resp
}
