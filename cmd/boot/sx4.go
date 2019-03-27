package main

import (
	"fmt"
	"xutil/cmd/boot/util"
)

func main() {

	files := util.GetFiles()

	for _, v := range files {

		fmt.Println(v)

	}

}
