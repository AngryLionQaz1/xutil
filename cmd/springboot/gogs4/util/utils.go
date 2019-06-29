package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

//移动文件
func Mv(path, s1, s2 string) {
	err := os.Rename(s1, s2)
	CheckErr(err, path)
}

//git 拉取文件
func Git(path, name, git string) {
	//1，判断文件夹是否存在
	if PathExists(filepath.Join(path, name)) {
		os.Chdir(filepath.Join(path, name))
		Exe(filepath.Join(path, name), "git", "pull")
	} else {
		os.Chdir(path)
		Exe(path, "git", "clone", git)
	}
}

//打印日志
func Logs(path, content string) {
	fd, _ := os.OpenFile(filepath.Join(path, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + content + "\n")
	defer fd.Close()

}

func CheckErr(err error, path string) {
	if err != nil {
		Logs(path, err.Error())
		log.Println("ERROR:%v", err)
	}

}

//获取系统类型
func GetOs() int {
	sysType := runtime.GOOS
	if sysType == "linux" {
		// LINUX系统
		return 1
	}
	if sysType == "windows" {
		// windows系统
		return 2
	}
	return 0
}

//判断文件或文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false

}

/**执行命令*/
func Exe(path, cm string, args ...string) {
	cmd := exec.Command(cm, args...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	CheckErr(err, path)
	//执行命令
	if err := cmd.Start(); err != nil {
		CheckErr(err, path)
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
				break
			}
			CheckErr(err, path)
			return
		}
		s := string(output)
		fmt.Printf("%s\n", s)
		Logs(path, s)
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		CheckErr(err, path)
		return
	}
}
