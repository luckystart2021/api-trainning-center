package utils

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func FileUpload(r *http.Request, fileName string) (string, error) {
	//ParseMultipartForm parses a request body as multipart/form-data
	file, handler, err := r.FormFile("img") //retrieve the file from form data
	//replace file with the key your sent your image with
	if err != nil {
		return "", err
	}
	checkFilenameExtension := FilenameExtension(handler.Filename)
	if checkFilenameExtension != ".jpg" || checkFilenameExtension != ".png" {
		return "", errors.New("File hình ảnh không hợp lệ")
	}
	defer file.Close() //close the file when we finish
	//this is path which  we want to store the file
	saveFileName := "upload/img/" + fileName + "/"
	// info, err := os.Stat()
	fileImg := FilenameWithoutExtension(handler.Filename) + buildFileName() + FilenameExtension(handler.Filename)
	f, err := os.OpenFile(saveFileName+fileImg, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)
	//here we save our file to our path
	return fileImg, nil
}

func buildFileName() string {
	return time.Now().Format("20060102150405")
}

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func FilenameExtension(fn string) string {
	return filepath.Ext(fn)
}
