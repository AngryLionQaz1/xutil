package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9999")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {

		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		s, err := ioutil.ReadAll(conn)
		fmt.Println(string(s))
		conn.Write([]byte(daytime))
		conn.Close()
	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
