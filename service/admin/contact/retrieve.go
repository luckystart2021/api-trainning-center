package contact

import (
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type Contact struct {
	FullName  string `json:"full_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	Subject   string `json:"subject"`
	CreatedAt string `json:"created_at"`
}

func (tc StoreContact) ShowContacts() ([]Contact, error) {
	contact, err := retrieveContact(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowContacts] error : ", err)
		return []Contact{}, err
	}

	return contact, nil
}

func retrieveContact(db *sql.DB) ([]Contact, error) {
	contacts := []Contact{}
	query := `
	SELECT 
		fullname, phone, email, message, subject, created_at
	FROM 
		contact
	ORDER BY created_at DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveContact] query error  %v", err)
		return contacts, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		var fullName, phone, message string
		var email, subject sql.NullString
		var createdAt time.Time
		err = rows.Scan(&fullName, &phone, &email, &message, &subject, &createdAt)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[retrieveContact] Scan error  %v", err)
			return contacts, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		contact := Contact{
			FullName:  fullName,
			Phone:     phone,
			Message:   message,
			CreatedAt: utils.TimeIn(createdAt, "VN", "02-01-2006 15:04:05"),
		}
		if email.Valid {
			contact.Email = email.String
		}
		if subject.Valid {
			contact.Subject = subject.String
		}
		contacts = append(contacts, contact)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveContact] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(contacts) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveContact] No Data  %v", err)
		return contacts, errors.New("Không có dữ liệu từ hệ thống")
	}
	return contacts, nil
}
