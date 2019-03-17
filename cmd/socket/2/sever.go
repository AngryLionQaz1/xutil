package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	server := ":8099"
	listener, err := net.Listen("tcp", server)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()
	readChan := make(chan string)
	writeChan := make(chan string)
	stopChan := make(chan bool)

	go readConn(conn, readChan, stopChan)
	go writeConn(conn, writeChan, stopChan)
	for {
		select {
		case readStr := <-readChan:
			upper := strings.ToUpper(readStr)
			writeChan <- upper
		case stop := <-stopChan:
			if stop {
				break
			}
		}
	}

}

func writeConn(conn net.Conn, writeChan chan string, stopChan chan bool) {

	for {
		strData := <-writeChan
		_, err := conn.Write([]byte(strData))
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println("发送:", strData)
	}

	stopChan <- true

}

func readConn(conn net.Conn, readChan chan string, stopChan chan bool) {

	for {
		data := make([]byte, 1024)
		_, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		strData := string(data)
		fmt.Println("接受到：", strData)
		readChan <- strData
	}
	stopChan <- true
}
