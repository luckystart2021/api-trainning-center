package album

import (
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/sirupsen/logrus"
)

func (st StoreAlbum) DeleteAlbum(id int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := deleteAlbumByRequest(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteAlbumByRequest] Delete album DB err  %v", err)
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa thông tin album thành công"
	} else {
		resp.Status = false
		resp.Message = "Xóa thông tin album không thành công"
	}
	return resp, nil
}

func deleteAlbumByRequest(db *sql.DB, id int) (int64, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	queryDeletePhotos := `
	DELETE 
		FROM 
			photos
		WHERE 
			id_album=$1;
	`
	_, err = tx.ExecContext(ctx, queryDeletePhotos, id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteAlbumByRequest] Delete photos DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	queryDeleteAlbum := `
	DELETE 
		FROM 
			album
		WHERE 
			id=$1;
	`
	_, err = tx.ExecContext(ctx, queryDeleteAlbum, id)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteAlbumByRequest] Delete album DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return 1, nil
}
