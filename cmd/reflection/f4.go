package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 123
	v := reflect.ValueOf(&x)
	//传递指针才能修改
	v.Elem().SetInt(999)
	fmt.Println(x)
}
