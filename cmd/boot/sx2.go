package main

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"os"
	"sort"
	_ "xutil/cmd/boot/resources"
)

func main() {

	sss, err := fs.New()
	if err != nil {
		panic(err)
		return
	}
	var files []string
	err = fs.Walk(sss, "/", func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
		return
	}
	sort.Strings(files)
	for _, v := range files {
		fmt.Println(v)
	}

}
