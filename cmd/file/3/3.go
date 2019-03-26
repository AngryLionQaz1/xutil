package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	buffer := bytes.NewBuffer([]byte{})
	f, err := os.OpenFile("sxs.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(buffer, f)
	io.WriteString(writer, "hello ")
	fmt.Println(buffer.String())
}
