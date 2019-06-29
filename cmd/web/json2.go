package main

import (
	"fmt"

	"github.com/json-iterator/go"
)

func main() {

	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	str := jsoniter.Get(val, "Colors", 0).ToString()
	i := jsoniter.Get(val, "ID").ToInt()
	fmt.Println(str)
	fmt.Println(i)
}
