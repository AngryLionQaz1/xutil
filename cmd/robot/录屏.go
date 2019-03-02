package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"strconv"
	"time"
)

func main() {

	for {
		bitmap := robotgo.CaptureScreen()
		fmt.Println("...", bitmap)
		time.Sleep(10 * time.Microsecond)

		robotgo.SaveBitmap(bitmap, "asses/img/"+strconv.FormatInt(time.Now().UnixNano(), 10)+".png")
		robotgo.FreeBitmap(bitmap)
	}

	//fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	//fmt.Printf("时间戳（纳秒）：%v;\n",strconv.FormatInt(time.Now().UnixNano(),10))
	//fmt.Printf("时间戳（毫秒）：%v;\n",time.Now().UnixNano() / 1e6)
	//fmt.Printf("时间戳（纳秒转换为秒）：%v;\n",time.Now().UnixNano() / 1e9)

}
