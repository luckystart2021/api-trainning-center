package album

import (
	"api-trainning-center/models/admin/photo"
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreAlbum) CreateAlbum(name, meta string) (response.MessageAlbumResponse, error) {
	response := response.MessageAlbumResponse{}
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateAlbum] err  %v", err)
		return response, errors.New("Không có dữ liệu từ hệ thống")
	}

	if err := CreateAlbumByRequest(ctx, tx, name, meta); err != nil {
		return response, err
	}

	album, err := selectAlbumByRequest(ctx, tx, name)
	if err != nil {
		return response, err
	}
	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateAlbum] err  %v", err)
		return response, errors.New("Không có dữ liệu từ hệ thống")
	}
	response.Id = album.Id
	response.Status = true
	response.Message = "Thêm album thành công"
	return response, nil
}

func selectAlbumByRequest(ctx context.Context, tx *sql.Tx, name string) (photo.Album, error) {
	album := photo.Album{}
	query := `
	SELECT
		id,
		name,
		meta
	FROM
		album
	WHERE
		name = $1;
	`
	rows := tx.QueryRowContext(ctx, query, name)
	err := rows.Scan(&album.Id, &album.Name, &album.Meta)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[selectAlbumByRequest] No Data  %v", err)
		tx.Rollback()
		return album, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		tx.Rollback()
		logrus.WithFields(logrus.Fields{}).Errorf("[selectAlbumByRequest] Scan error  %v", err)
	}
	return album, nil
}

func CreateAlbumByRequest(ctx context.Context, tx *sql.Tx, name, meta string) error {
	query := `
	INSERT INTO album (name, meta)
	VALUES($1,$2);
	`
	_, err := tx.ExecContext(ctx, query, name, meta)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateAlbumByRequest] Insert album DB err  %v", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
