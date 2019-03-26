package main

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"log"
	"os"
	_ "xutil/statik"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	f, _ := statikFS.Open("/gos.go")

	fmt.Println(f.Stat())

}

//读取文件到[]byte中
func file2bytes(filename string) ([]byte, error) {

	//File
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	//FileInfo
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	//[]byte
	data := make([]byte, stats.Size())
	count, err := file.Read(data)

	if err != nil {
		return nil, err
	}
	fmt.Printf("read file %s len: %d \n", filename, count)

	return data, nil

}
