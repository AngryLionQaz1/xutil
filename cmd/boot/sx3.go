package main

import (
	"fmt"
	"xutil/cmd/boot/util"
)

func main() {

	err := util.WriteToFile("boot/pom.xml", []byte("ssssss"))
	fmt.Println(err)
}
