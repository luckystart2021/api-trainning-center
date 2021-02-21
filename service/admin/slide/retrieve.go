package slide

import (
	"api-trainning-center/models/admin/slide"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type ShowSlides struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

func (tc StoreSlide) ShowSlides() ([]ShowSlides, error) {
	responseSlides := []ShowSlides{}
	slides, err := retrieveSlide(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowSlides] error : ", err)
		return []ShowSlides{}, err
	}
	for _, data := range slides {
		if data.Hide == true {
			continue
		}
		slide := ShowSlides{
			Title: data.Title,
			Img:   "/files/img/slide/" + data.Img,
		}
		responseSlides = append(responseSlides, slide)
	}
	if len(responseSlides) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveSlide] No Data  %v", err)
		return responseSlides, errors.New("Không có dữ liệu từ hệ thống")
	}

	return responseSlides, nil
}

func (tc StoreSlide) ShowSlidesAdmin() ([]slide.Slide, error) {
	slides, err := retrieveSlide(tc.db)
	if len(slides) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[ShowSlidesAdmin] No Data  %v", err)
		return []slide.Slide{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowSlidesAdmin] error : ", err)
		return []slide.Slide{}, err
	}
	return slides, nil
}

func (tc StoreSlide) ShowDetailSlide(idSlide int) (slide.Slide, error) {
	detailSlide, err := retrieveDetailSlide(tc.db, idSlide)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowSlidesAdmin] error : ", err)
		return slide.Slide{}, err
	}
	return detailSlide, nil
}

func retrieveDetailSlide(db *sql.DB, idSlide int) (slide.Slide, error) {
	slide := slide.Slide{}
	query := `
	SELECT
		id,
		title,
		img,
		hide,
		created_at,
		created_by
	FROM
		slide
	WHERE 
		id = $1;
	`
	rows := db.QueryRow(query, idSlide)
	var createdAt time.Time
	err := rows.Scan(&slide.Id, &slide.Title, &slide.Img, &slide.Hide, &createdAt, &slide.CreatedBy)
	slide.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveDetailSlide] No Data  %v", err)
		return slide, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveDetailSlide] Scan error  %v", err)
		return slide, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return slide, nil
}

func retrieveSlide(db *sql.DB) ([]slide.Slide, error) {
	slides := []slide.Slide{}
	query := `
	SELECT
		id,
		title,
		img,
		hide,
		created_at,
		created_by
	FROM
		slide;
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveSlide] query error  %v", err)
		return slides, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	for rows.Next() {
		var title, img, createBy string
		var createdAt time.Time
		var id int
		var hide bool
		err = rows.Scan(&id, &title, &img, &hide, &createdAt, &createBy)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveSlide] Scan error  %v", err)
			return slides, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		slide := slide.Slide{
			Id:        id,
			Title:     title,
			Img:       img,
			Hide:      hide,
			CreatedBy: createBy,
			CreatedAt: utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS),
		}
		slides = append(slides, slide)
	}
	if len(slides) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveSlide] No Data  %v", err)
		return slides, errors.New("Không có dữ liệu từ hệ thống")
	}
	return slides, nil
}
