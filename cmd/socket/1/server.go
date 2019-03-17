package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	server := ":8099"
	netListen, err := net.Listen("tcp", server)
	defer netListen.Close()
	if err != nil {
		log("链接错误", err)
		os.Exit(1)
	}
	log("等待客户端链接")
	for {

		conn, err := netListen.Accept()
		if err != nil {
			log(conn.RemoteAddr().String())
			continue
		}
		//设置短链接（10s）
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
		log(conn.RemoteAddr().String(), "链接成功")
		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log(conn.RemoteAddr().String())
			return
		}
		Data := buffer[:n]
		message := make(chan byte)
		//心跳计时
		go heartBeating(conn, message, 4)
		//检测每次是否有数据传入
		go gravelChannel(Data, message)
		log(time.Now().Format("2006-01-02 15:04:05.0000000"), conn.RemoteAddr().String(), string(buffer[:n]))

	}

	defer conn.Close()
}

func gravelChannel(bytes []byte, mess chan byte) {
	for _, v := range bytes {
		mess <- v
	}
	close(mess)
}

func heartBeating(conn net.Conn, bytes chan byte, timeout int) {
	select {
	case fk := <-bytes:
		log(conn.RemoteAddr().String(), "心跳:第", string(fk), "times")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(5 * time.Second):
		log("超时关闭链接")
		conn.Close()
	}
}

func log(v ...interface{}) {
	fmt.Println(v...)
	return
}
