package main

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"log"
	_ "xutil/cmd/boot/resources"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, _ := statikFS.Open("/src/main/resources/config/application-actuator.yml")
	fmt.Println(f.Stat())

}
