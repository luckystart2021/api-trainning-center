package information

import (
	"database/sql"
	"errors"
)

func CountInformation(db *sql.DB) (int, error) {
	var count int
	query := `
	SELECT COUNT(*) FROM information;
	`
	row := db.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return count, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return count, nil
}
