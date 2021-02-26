package upload

import (
	"api-trainning-center/service/response"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func Upload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reps := response.UploadResponse{}
		saveFileName := "upload/img/ckeditor" + "/"
		if _, err := os.Stat(saveFileName); os.IsNotExist(err) {
			os.Mkdir(saveFileName, 0755)
		}
		fileImg := buildFileName() + ".jpg"
		file, err := os.OpenFile(saveFileName+fileImg, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[UploadSigleIMG] request: ", err)
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Lỗi hệ thống khi upload hình ảnh"))
			return
		}
		_, err = io.Copy(file, r.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[UploadSigleIMG] request: ", err)
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Lỗi hệ thống khi upload hình ảnh"))
			return
		}
		reps.FileName = "/files/img/ckeditor/" + fileImg
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, reps)
	}
}

func buildFileName() string {
	return time.Now().Format("20060102150405")
}
