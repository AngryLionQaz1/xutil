package prt

import "fmt"

/***
打印信息
*/

const (
	qback    = 30
	qred     = 31
	qgreen   = 32
	qyellow  = 33
	qblue    = 34
	qpurple  = 35
	qcyanine = 36
	qwhite   = 37
	bback    = 40
	bred     = 41
	bgreen   = 42
	byellow  = 43
	bblue    = 44
	bpurple  = 45
	bcyanine = 46
	bwhite   = 47
	x0       = 0 // 终端默认设置
	x1       = 1 // 高亮显示
	x4       = 4 //使用下划线
	x5       = 5 // 闪烁
	x7       = 7 //反白显示
	x8       = 8 //不可见
)

func PrintlnErr(err error) {
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, x1, 0, qred, err, 0x1B)
}

func PrintlnRight(str string) {
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, x1, 0, qgreen, str, 0x1B)
}

func PrintlnTips(str string) {
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, x1, 0, qyellow, str, 0x1B)
}
