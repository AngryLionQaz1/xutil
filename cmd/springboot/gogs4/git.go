package main

import (
	"fmt"
	"xutil/cmd/springboot/gogs4/jar"
)

func main() {

	//util.Git(`D:\`,"z","xiaoyi@localhost:xiaoyiqaz1/z.git")
	//util.Exe(`D:\`, "git", "pull")

	path := jar.GetPath("c,x")

	fmt.Println(path)
}
