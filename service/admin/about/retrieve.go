package about

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreAbout) ShowAbout() ([]About, error) {
	showAbout, err := retrieveAbout(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowAbout] error : ", err)
		return []About{}, err
	}
	return showAbout, nil
}

func retrieveAbout(db *sql.DB) ([]About, error) {
	abouts := []About{}
	query := `
	SELECT
		title,
		description,
		subtitle,
		img
	FROM 
		notification
	ORDER BY 
		id;
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAbout] query error  %v", err)
		return abouts, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	for rows.Next() {
		var title, description, subtitle, img string
		err = rows.Scan(&title, &description, &subtitle, &img)

		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAbout] Scan error  %v", err)
			return abouts, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		about := About{
			Title:       title,
			Description: description,
			SubTitle:    subtitle,
			Img:         "/files/img/about/" + img,
		}

		abouts = append(abouts, about)
	}

	if len(abouts) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveAbout] No Data  %v", err)
		return abouts, errors.New("Không có dữ liệu từ hệ thống")
	}

	return abouts, nil
}
