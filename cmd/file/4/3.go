package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	fw, err := os.Create("demo.gzip") // 创建gzip包文件，返回*io.Writer
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	// 实例化心得gzip.Writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 获取要打包的文件信息
	fr, err := os.Open(`E:\GoProjects\xutil\asses\gos.go`)
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 获取文件头信息
	fi, err := fr.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// 创建gzip.Header
	gw.Header.Name = fi.Name()

	// 读取文件数据
	buf := make([]byte, fi.Size())
	_, err = fr.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// 写入数据到zip包
	_, err = gw.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
}
