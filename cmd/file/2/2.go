package main

import (
	"bytes"
	"fmt"
)

var b bytes.Buffer

func main() {

	b := bytes.NewBufferString("s")
	fmt.Println(b)
}
