package main

import (
	"fmt"
	"syscall"
)

func main() {

	WindowVersion1()

}

func WindowVersion1() {
	h, err := syscall.LoadLibrary("dxgi.dll")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	print(h)
}
