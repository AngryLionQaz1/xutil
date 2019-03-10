package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	dir, _ := os.Getwd()
	fmt.Println(dir)
	os.Chdir("C:\\Users\\pc\\Desktop\\2w")
	fmt.Println(os.Getwd())
	osPath, _ := exec.LookPath(os.Args[0])
	fmt.Println(osPath)
	dir2, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir2)

	join := filepath.Join(dir2, "sx", "init.html")

	fmt.Println(join)

}
