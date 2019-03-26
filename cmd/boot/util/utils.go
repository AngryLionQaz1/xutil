package util

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"log"
	"os"
	"sort"
)

//获取文件
func GetFiles() []string {
	sss, err := fs.New()
	CheckError(err)
	var files []string
	err = fs.Walk(sss, "/", func(path string, fi os.FileInfo, err error) error {
		CheckError(err)
		files = append(files, path)
		return nil
	})
	CheckError(err)
	sort.Strings(files)
	for _, v := range files {
		fmt.Println(v)
	}
	return files
}

//

//创建目录
func CreateDir(path string) {
	//创建多级目录和设置权限
	os.MkdirAll(path, 0777)
}

//创建文件
func CreateFile(path, content string) {
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err == nil {
		//fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
	}
}

//判断是文件还是文件夹
func CheckDir(f *os.File) bool {
	fs, err := f.Stat()
	CheckError(err)
	return fs.IsDir()
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
