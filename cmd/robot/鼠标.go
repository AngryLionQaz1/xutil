package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {

	robotgo.ScrollMouse(10, "up")
	//robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 0.1, 1.0)
	//robotgo.MoveMouse(100,100)
	robotgo.DragMouse(100, 100)
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

}
