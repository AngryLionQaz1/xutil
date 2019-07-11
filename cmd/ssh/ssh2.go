package main

import (
	"fmt"
	"net/url"
	"xutil/cmd/ssh/util"
)

func main() {

	client := util.SSHClient{}
	u := &url.URL{Scheme: "ssh", Host: "192.168.3.53:22"}
	channel := make(chan []byte, 10)
	client.Command(*u, "root", "xiaoyi", "pwd", channel)

	fmt.Println(string(<-channel))
}
