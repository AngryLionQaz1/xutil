package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {

	for {
		bitmap := robotgo.CaptureScreen()
		fmt.Println("...", bitmap)
		time.Sleep(100 * time.Microsecond)
		robotgo.FreeBitmap(bitmap)
	}

}
