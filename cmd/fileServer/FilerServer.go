package main

import (
	"fmt"
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
	"github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//http 服务
var server *http.Server
var quit = make(chan os.Signal)

const (
	initFile = "init.yaml"
)

type Program struct {
	//服务名
	Name string `yaml:name`
	//服务器显示
	DisplayName string `yaml:displayName`
	//服务描述
	Description string `yaml:description`
	//路径
	Path string `yaml:path`
	//端口号
	Port int `yaml:port`
	//文件大小
	Size int `yaml:size`
}

func main() {

	program := initProgram()
	config := initService(program)
	s, err := service.New(program, config)
	checkErr(err)
	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}
	cmd := os.Args[1]
	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	}

	if cmd == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	}

}

func (p *Program) run() {

	fileServer(p.Path, p.Port, p.Size)

}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	logs("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	logs("停止服务")
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal([]byte("shutdown the server err"))
	}
	return nil
}

//初始化服务
func initService(p *Program) *service.Config {

	var serviceConfig = &service.Config{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}
	return serviceConfig
}

//读取配置文件
func initProgram() *Program {
	i := new(Program)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, initFile))
	checkErr(err)
	err = yaml.Unmarshal(yamlFile, i)
	checkErr(err)
	return i
}

//文件服务
func fileServer(dir string, port, size int) {

	signal.Notify(quit, os.Interrupt)
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(dir))
	mux.Handle("dir", fileServer)
	mux.Handle("/upload", uploadHandler(dir, size))
	mux.HandleFunc("/exit", bye)
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if filter(request.URL) {
			writer.WriteHeader(http.StatusNotFound)
			io.WriteString(writer, "404 page not found!\n")
			return
		}
		fileServer.ServeHTTP(writer, request)
	})
	server = &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}
	go func() {
		//接受退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}

		log.Fatal("Server exited")
	}

}

func uploadHandler(path string, size int) http.Handler {

	maxUploadSize := int64(size * 1024 * 1024) //M

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}
		fileType := r.PostFormValue("type")
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
		fileType = http.DetectContentType(fileBytes)
		if fileType != "image/jpeg" && fileType != "image/jpg" &&
			fileType != "image/gif" && fileType != "image/png" &&
			fileType != "application/pdf" {
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		v4 := uuid.NewV4()
		fmt.Printf("UUIDv4: %s\n", v4)

		fileName := v4.String()
		fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}

		newPath := filepath.Join(filepath.Join(CreateDateDir(path), fileName), fileName+fileEndings[0])
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

//关闭http
func bye(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("bye bye ,shutdown the server")) // 没有输出
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal([]byte("shutdown the server err"))
	}
}

//拦截文件夹
func filter(url *url.URL) bool {
	s := url.Path
	split := strings.Split(s, ".")
	if len(split) == 2 {
		return false
	}
	return true

}
func checkErr(err error) {

	if err != nil {
		log.Printf("ERROR: #%v ", err)
		logs(err.Error())
	}

}

//打印日志
func logs(s string) {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fd, _ := os.OpenFile(filepath.Join(dir, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s + "\n")
	defer fd.Close()

}
