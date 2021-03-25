package upload

import (
	"api-trainning-center/service/constant"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func buildFileName() string {
	return time.Now().Format("20060102150405")
}

//CkUpload handles /ckupload route
func CkUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reps := response.UploadResponse{}
		if r.Method == "POST" {
			err := r.ParseMultipartForm(32 << 20)
			if err != nil {
				log.Printf("ERROR: %s\n", err)
				http.Error(w, err.Error(), 500)
				return
			}
			mpartFile, mpartHeader, err := r.FormFile("upload")
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer mpartFile.Close()
			uri, err := saveFile(mpartHeader, mpartFile)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// CKEdFunc := r.FormValue("CKEditorFuncNum")
			// fmt.Fprintln(w, "<script>window.parent.CKEDITOR.tools.callFunction("+CKEdFunc+", \""+uri+"\");</script>")
			reps.Url = constant.Domain + "/files/img/ck/" + uri
			response.RespondWithJSON(w, http.StatusOK, reps)
		} else {
			err := fmt.Errorf("Method %q not allowed", r.Method)
			log.Printf("ERROR: %s\n", err)
			http.Error(w, err.Error(), 405)
		}
	}
}

//saveFile saves file to disc and returns its relative uri
func saveFile(fh *multipart.FileHeader, f multipart.File) (string, error) {
	saveFileName := "upload/img/ck/"
	if _, err := os.Stat(saveFileName); os.IsNotExist(err) {
		os.Mkdir(saveFileName, 0755)
	}
	// info, err := os.Stat()
	fileImg := utils.FilenameWithoutExtension(fh.Filename) + "_" + buildFileName() + utils.FilenameExtension(fh.Filename)
	file, err := os.OpenFile(saveFileName+fileImg, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, f)
	if err != nil {
		return "", err
	}
	return fileImg, nil
}
