package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const paths = `E:\GoProjects\xutil`

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/up", func(writer http.ResponseWriter, request *http.Request) {

		//headers := request.Header
		request.ParseMultipartForm(-1)
		files := request.MultipartForm.File["file"]
		for i, h := range files {
			fmt.Println(i, h.Filename)

			checkFileType(path.Ext(h.Filename))

		}

	})

	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		maxUploadSize := int64(1000 * 1024 * 1024) //M
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}
		files := r.MultipartForm.File["file"]
		fmt.Println(path.Ext(files[0].Filename))
		file, _, err := r.FormFile("file")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		fileType := http.DetectContentType(fileBytes)

		if fileType != "image/jpeg" && fileType != "image/jpg" &&
			fileType != "image/gif" && fileType != "image/png" &&
			fileType != "application/pdf" {
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		v4 := strings.Replace(uuid.NewV4().String(), "-", "", -1)
		fmt.Printf("UUIDv4: %s\n", v4)
		fileName := v4
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		fileEndings, err := mime.ExtensionsByType(fileType)
		newPath := filepath.Join(filepath.Join(CreateDateDir(paths), fileName+fileEndings[0]))
		fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()
		if _, err := newFile.Write(fileBytes); err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))

	})

	http.ListenAndServe(":8086", mux)

}

func renderError(writer http.ResponseWriter, s string, i int) {
	writer.WriteHeader(i)
	io.WriteString(writer, s)
}

// CreateDateDir 根据当前日期来创建文件夹
func CreateDateDir(basePath string) string {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

func checkFileType(s string) bool {

	str := `jar,bat`
	fmt.Println(strings.Contains(str, strings.Replace(s, ".", "", -1)))

	return true

}
