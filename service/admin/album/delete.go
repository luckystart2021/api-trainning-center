package album

import (
	photoService "api-trainning-center/service/admin/photo"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/sirupsen/logrus"
)

func (st StoreAlbum) DeleteAlbum(id int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	photos, err := photoService.FindPhotosByIdAlbum(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindPhotosByIdAlbum] error : ", err)
		return resp, err
	}

	count, err := deleteAlbumByRequest(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteAlbumByRequest] Delete album DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		for _, data := range photos {
			err = utils.DeleteFile("upload/img/album/" + data.Img)
			if err != nil {
				logrus.WithFields(logrus.Fields{}).Error("[DeletePhoto] error : ", err)
				return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
			}
		}
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
