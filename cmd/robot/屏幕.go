package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)
	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color----", color)

	fmt.Println(robotgo.GetScreenSize())
	scree := robotgo.CaptureScreen()
	fmt.Println(scree)

}
