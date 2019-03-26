package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	// 打开gzip文件
	fr, err := os.Open("demo.gzip")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	// 创建gzip.Reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		log.Fatalln(err)
	}
	defer gr.Close()

	// 读取文件内容
	buf := make([]byte, 1024*1024*10) // 如果单独使用，需自己决定要读多少内容，根据官方文档的说法，你读出的内容可能超出你的所需（当你压缩gzip文件中有多个文件时，强烈建议直接和tar组合使用）
	n, err := gr.Read(buf)

	// 将包中的文件数据写入
	fw, err := os.Create(gr.Header.Name)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fw.Write(buf[:n])
	if err != nil {
		log.Fatalln(err)
	}
}
