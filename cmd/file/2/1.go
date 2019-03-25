package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"xutil/asses"
)

func main() {

	data, _ := file2bytes(`/home/xiaoyi/project/golandProjects/xutil/asses/sss.xlsx`)

	data2 := `package asses

const Gos ="` + bytesToHexString(data) + `" 
`

	gos("gos.go", []byte(data2))

	fmt.Printf(asses.Gos)

	gos("SX.xlsx", hexStringToBytes(asses.Gos))

}

//生成go文件
func gos(filename string, bytes []byte) {

	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	fd.Write(bytes)

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

// byte数组转16进制字符串
func bytesToHexString(b []byte) string {
	var buf bytes.Buffer
	for _, v := range b {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) > 1 {
			buf.WriteString(t)
		} else {
			buf.WriteString("0" + t)
		}
	}
	return buf.String()
}

// 16进制字符串转byte数组
func hexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}
