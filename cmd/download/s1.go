package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 主动关闭服务器
var server *http.Server

type fileHanlder struct {
}

func main() {

	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/s", &myHandler{})
	mux.HandleFunc("/bye", sayBye)
	mux.Handle("/files", &fileHanlder{})

	server = &http.Server{
		Addr:         ":8086",
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting v3 httpserver")
	err := server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("E:\\GoProjects\\xutil\\cmd\\download\\s1.go")
	checkErr(err)
	infor, _ := file.Stat()
	byteSlice := make([]byte, infor.Size())
	w.Write(byteSlice)

}

func (*fileHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/files", http.FileServer(http.Dir("C:\\Users\\win\\Desktop\\thunder"))).ServeHTTP(w, r)
}

// 关闭http
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the server")) // 没有输出
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal([]byte("shutdown the server err"))
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
