package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9999")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	conn.Write([]byte("HEAD /HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
