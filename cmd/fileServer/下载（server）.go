package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

var dir string
var port int
var staticHandler http.Handler

// 初始化参数
func init() {
	dir = path.Dir("C:\\Users\\win\\Desktop\\thunder")
	flag.IntVar(&port, "port", 8086, "服务器端口")
	flag.Parse()
	fmt.Println("dir:", http.Dir(dir))
	staticHandler = http.FileServer(http.Dir(dir))
}

func main() {
	http.HandleFunc("/", StaticServer)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("path:" + req.URL.Path)
	if req.URL.Path != "/down/" {
		staticHandler.ServeHTTP(w, req)
		return
	}

	io.WriteString(w, "hello, world!\n")
}
