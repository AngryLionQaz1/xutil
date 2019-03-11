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

func main() {

	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler3{})
	mux.HandleFunc("/bye", sayBye3)

	server = &http.Server{
		Addr:         ":8086",
		WriteTimeout: time.Second * 4,
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

type myHandler3 struct{}

func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}

// 关闭http
func sayBye3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown the server")) // 没有输出
	err := server.Shutdown(nil)
	if err != nil {
		log.Fatal([]byte("shutdown the server err"))
	}
}
