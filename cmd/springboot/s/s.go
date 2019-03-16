package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const jar = "file-0.0.1-SNAPSHOT.jar"

func main() {

	s := `
   pid=` + "`" + `ps -ef|grep ` + jar + `|grep -v grep|awk '{print $2}'` + "`" + `
  #如果不存在返回1，存在返回0
  if [ -z "${pid}" ]; then
    echo ">>> ` + jar + ` is not running <<<"
  else
    echo ">>> ` + jar + ` is running PID is ${pid} <<<"
  fi
`

	runSh(s)

}

//运行sh脚本
func runSh(sh string) {
	os.Chmod(sh, 0777)
	cmd := exec.Command("/bin/bash", "-c", sh)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	checkErr(err)
	//执行命令
	if err := cmd.Start(); err != nil {
		checkErr(err)
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				checkErr(err)
			}
			return
		}
		fmt.Printf("%s\n", string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
