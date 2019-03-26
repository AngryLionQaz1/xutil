package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	err := Zip("demo.zip", []string{`E:\GoProjects\xutil\asses\gos.go`})
	if err != nil {
		fmt.Println(err)
	}
}

func Zip(zipFile string, fileList []string) error {
	// 创建 zip 包文件
	fw, err := os.Create(zipFile)
	if err != nil {
		log.Fatal()
	}
	defer fw.Close()

	// 实例化新的 zip.Writer
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for _, fileName := range fileList {
		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}
		fi, err := fr.Stat()
		if err != nil {
			return err
		}
		// 写入文件的头信息
		fh, err := zip.FileInfoHeader(fi)
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return err
		}
		// 写入文件内容
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
	}
	return nil
}
