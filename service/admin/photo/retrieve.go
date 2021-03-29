package photo

import (
	"api-trainning-center/models/admin/photo"
	"api-trainning-center/utils"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
)

type PhotoResponse struct {
	Img   string `json:"img"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

type PhotosResponse struct {
	Id        int           `json:"id"`
	AlbumName string        `json:"album_name"`
	AlbumMeta string        `json:"album_meta"`
	Photos    PhotoResponse `json:"photo"`
}

func (st StorePhoto) ShowPhoto(idAlbum int) ([]PhotoResponse, error) {
	ps := []PhotoResponse{}
	photos, err := FindPhotosByIdAlbum(st.db, idAlbum)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindPhotosByIdAlbum] error : ", err)
		return nil, err
	}
	for _, data := range photos {
		p := PhotoResponse{}
		p.Img = "/files/img/album/" + data.Img
		p.Meta = data.Meta
		p.Title = data.Title
		ps = append(ps, p)
	}
	return ps, nil
}

func (st StorePhoto) ShowPhotoInAdmin(id int) (photo.PhotoResponse, error) {
	p := photo.PhotoResponse{}
	photo, err := FindOnePhoto(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowPhotoInAdmin] error : ", err)
		return p, err
	}
	p.Id = photo.Id
	p.IdAlbum = photo.IdAlbum
	p.Title = photo.Title
	p.Meta = photo.Meta
	p.Img = "/files/img/album/" + photo.Img
	p.CreatedBy = photo.CreatedBy
	p.UpdatedBy = photo.UpdatedBy
	p.CreatedAt = utils.TimeIn(photo.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	p.UpdatedAt = utils.TimeIn(photo.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	return p, nil
}

func FindOnePhoto(db *sql.DB, id int) (photo.Photo, error) {
	photo := photo.Photo{}
	query := `
	SELECT
		id,
		img,
		title,
		meta,
		created_by,
		created_at,
		updated_by,
		updated_at,
		id_album
	FROM
		photos
	WHERE
		id = $1;
	`
	rows := db.QueryRow(query, id)
	err := rows.Scan(&photo.Id, &photo.Img, &photo.Title, &photo.Meta, &photo.CreatedBy, &photo.CreatedAt, &photo.UpdatedBy, &photo.UpdatedAt, &photo.IdAlbum)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOnePhoto] No Data  %v", err)
		return photo, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOnePhoto] Scan error  %v", err)
	}
	return photo, nil
}

func (st StorePhoto) ShowPhotosInAdmin() ([]photo.Photo, error) {
	photos, err := FindAllPhotos(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindAllPhotos] error : ", err)
		return nil, err
	}
	return photos, nil
}

func FindAllPhotos(db *sql.DB) ([]photo.Photo, error) {
	photos := []photo.Photo{}
	query := `
	SELECT
		id,
		id_album,
		img,
		title,
		meta,
		created_by,
		created_at,
		updated_by,
		updated_at
	FROM
		photos
	ORDER BY id;
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllPhotos] query error  %v", err)
		return photos, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		photo := photo.Photo{}
		var title, meta null.String
		var img string
		err = rows.Scan(&photo.Id, &photo.IdAlbum, &img, &title, &meta, &photo.CreatedBy, &photo.CreatedAt, &photo.UpdatedBy, &photo.UpdatedAt)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[FindAllPhotos] Scan error  %v", err)
			return photos, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		if title.Valid {
			photo.Title = title.String
		}
		if meta.Valid {
			photo.Meta = meta.String
		}
		photo.Img = "/files/img/album/" + img
		photos = append(photos, photo)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllPhotos] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(photos) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllPhotos] No Data  %v", err)
		return photos, errors.New("Không có dữ liệu từ hệ thống")
	}
	return photos, nil
}

func (st StorePhoto) ShowPhotos() ([]PhotosResponse, error) {
	album, err := FindAlbum(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findAlbum] error : ", err)
		return nil, err
	}
	photosResp := []PhotosResponse{}
	for _, dataAlbum := range album {
		photosResponse := PhotosResponse{}
		photosResponse.Id = dataAlbum.Id
		photosResponse.AlbumName = dataAlbum.Name
		photosResponse.AlbumMeta = dataAlbum.Meta
		photos, err := findPhotosThumbnailByIdAlbum(st.db, dataAlbum.Id)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[findPhotosThumbnailByIdAlbum] error : ", err)
			continue
		}
		photosResponse.Photos.Img = "/files/img/album/" + photos.Img
		photosResponse.Photos.Meta = photos.Meta
		photosResponse.Photos.Title = photos.Title
		photosResp = append(photosResp, photosResponse)
	}
	return photosResp, nil
}

func findPhotosThumbnailByIdAlbum(db *sql.DB, id int) (photo.Photo, error) {
	photo := photo.Photo{}
	query := `
	SELECT
		img,
		title,
		p.meta
	FROM
		photos p
	WHERE id_album = $1
	ORDER BY created_at DESC
	LIMIT 1;
	`
	rows := db.QueryRow(query, id)
	err := rows.Scan(&photo.Img, &photo.Title, &photo.Meta)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosThumbnailByIdAlbum] No Data  %v", err)
		return photo, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosThumbnailByIdAlbum] Scan error  %v", err)
		return photo, errors.New("Không có dữ liệu từ hệ thống")
	}
	return photo, nil
}

func FindAlbum(db *sql.DB) ([]photo.Album, error) {
	albums := []photo.Album{}
	query := `
	SELECT
		id,
		name,
		meta
	FROM
		album
	ORDER BY id;
	`
	rows, err := db.Query(query)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] query error  %v", err)
		return albums, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		album := photo.Album{}
		err = rows.Scan(&album.Id, &album.Name, &album.Meta)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] Scan error  %v", err)
			return albums, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		albums = append(albums, album)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(albums) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] No Data  %v", err)
		return albums, errors.New("Không có dữ liệu từ hệ thống")
	}
	return albums, nil
}

func FindPhotosByIdAlbum(db *sql.DB, idAlbum int) ([]photo.Photo, error) {
	photos := []photo.Photo{}
	query := `
	SELECT
		p.id,
		img,
		title,
		p.meta,
		created_by,
		created_at,
		updated_by,
		updated_at,
		id_album
	FROM
		photos p
	INNER JOIN 
		album a ON a.id = p.id_album
	WHERE
		id_album = $1;
	`
	rows, err := db.Query(query, idAlbum)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosByIdAlbum] query error  %v", err)
		return photos, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	defer rows.Close()
	for rows.Next() {
		photo := photo.Photo{}
		err = rows.Scan(&photo.Id, &photo.Img, &photo.Title, &photo.Meta, &photo.CreatedBy, &photo.CreatedAt, &photo.UpdatedBy, &photo.UpdatedAt, &photo.IdAlbum)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosByIdAlbum] Scan error  %v", err)
			return photos, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		photos = append(photos, photo)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosByIdAlbum] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return photos, nil
}
