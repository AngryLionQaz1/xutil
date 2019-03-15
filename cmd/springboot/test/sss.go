package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {

	exe("/bin/bash", "gogs", "-web", "&")

}
func exe(cm string, args ...string) {

	fd, _ := os.OpenFile("gogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	cmd := exec.Command(cm, args...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	checkErr(err)
	//执行命令
	if err := cmd.Start(); err != nil {
		checkErr(err)
		return
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		//_, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				//fd.Close()
				checkErr(err)
			}
			return
		}
		fd.WriteString("\n")
		fd.WriteString(string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		checkErr(err)
		return
	}
}

func checkErr(err error) {

	if err != nil {
		log.Printf("ERROR: #%v ", err)

	}

}
