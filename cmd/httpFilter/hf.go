package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"xutil/cmd/httpFilter/web"
)

type HttpServer struct {
	http.Server
}

func (server *HttpServer) StartServer() {
	log.Println("web server start " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (server *HttpServer) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("测试")
}

func TestWebFilter() {
	web.Register("/safe/**", func(rw http.ResponseWriter, r *http.Request) error {
		return errors.New("解密失败")
		//return nil
	})
	web.Register("/safe/user/**", func(rw http.ResponseWriter, r *http.Request) error {
		return errors.New("请登录")
		//return nil
	})
	http.HandleFunc("/safe", web.Handle(func(wr http.ResponseWriter, req *http.Request) error {
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	http.HandleFunc("/safe/user/test", web.Handle(func(wr http.ResponseWriter, req *http.Request) error {
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	http.HandleFunc("/safe/user", web.Handle(func(wr http.ResponseWriter, req *http.Request) error {
		wr.Write([]byte(req.RequestURI))
		return nil
	}))

	server := &HttpServer{}
	server.Addr = ":8080"
	server.StartServer()
}

func main() {
	TestWebFilter()
}
