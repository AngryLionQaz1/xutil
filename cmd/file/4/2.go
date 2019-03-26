package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	err := Unzip("demo.zip")
	if err != nil {
		fmt.Println(err)
	}
}

func Unzip(zipFile string) error {
	zr, err := zip.OpenReader(zipFile)
	defer zr.Close()
	if err != nil {
		return err
	}

	for _, file := range zr.File {
		// 如果是目录，则创建目录
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(file.Name, file.Mode()); err != nil {
				return err
			}
			continue
		}
		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		fw, err := os.OpenFile(file.Name, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, fr)
		if err != nil {
			return err
		}
		fw.Close()
		fr.Close()
	}
	return nil
}
