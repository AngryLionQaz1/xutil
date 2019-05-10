package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	//获取当前文件路径
	s, _ := os.Getwd()
	fmt.Println(s)
	Start("auth-0.0.1-SNAPSHOT.jar")

}

//运行jar程序
func Start(jar string) {
	//cmd := exec.Command( "nohup","java","-Xms256m","-Xmx512m","-jar",jar,"/dev/null","2",">","&1","&")
	cmd := exec.Command("java", "-Xms256m", "-Xmx512m", "-jar", jar)
	err := cmd.Start()
	checkErr(err)
	fmt.Println(cmd.Process.Pid)
	wPid(cmd.Process.Pid, jar)
}
func checkErr(err error) {
	if err != nil {
		log.Println("ERROR: #%v ", err)
		logs(err.Error())
	}
}

//打印日志
func logs(s string) {
	fd, _ := os.OpenFile(filepath.Join("logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s + "\n")
	defer fd.Close()
}

//记录pid
func wPid(pid int, jar string) {
	s, _ := os.Getwd()
	fd, _ := os.OpenFile(filepath.Join(s, jar+".pid"), os.O_RDWR|os.O_CREATE, 0644)
	fd.WriteString(strconv.Itoa(pid))
	defer fd.Close()
}
