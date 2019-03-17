package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {

	execCommand("git", "clone", "root@47.92.213.93:xiaoyiqaz1/file.git")
	execCommand("git", "pull", "origin", "master")

}

//执行命令函数
//commandName 命名名称，如cat，ls，git等
//params 命令参数，如ls -l的-l，git log 的log等
func execCommand(commandName string, params ...string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
		logs(line)
	}

	cmd.Wait()
	return true
}

//打印日志
func logs(s string) {

	fd, _ := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s)
	defer fd.Close()

}
