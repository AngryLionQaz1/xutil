package main

import (
	"fmt"
	"syscall"
)

func main() {

	exportgo := syscall.NewLazyDLL(`E:\GoProjects\xutil\cmd\exportgo.dll`)

	proc := exportgo.NewProc("Sum")

	ret, _, _ := proc.Call(1, 3, 3)
	fmt.Println(ret)
}
