package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {
	//bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
	bitmap := robotgo.CaptureScreen()
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)
	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	fmt.Println("FindBitmap------", fx, fy)

	//robotgo.SaveBitmap(bitmap, string(time.Now().UnixNano())+".png")
	robotgo.SaveBitmap(bitmap, "asses/img/"+"x.png")

	//robotgo.Convert("test.png", "test.tif")

	//portion := robotgo.GetPortion(bitmap, 100, 100, 200, 200)
	//
	//robotgo.SaveBitmap(portion, "portion.png")

}
