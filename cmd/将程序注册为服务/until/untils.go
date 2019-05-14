package until

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//运行程序
func Start(name, path, cmd, args string) {
	strs := []string{}
	for _, v := range strings.Split(args, ",") {
		strs = append(strs, v)
	}
	command := exec.Command(cmd, strs...)
	command.Start()
	fmt.Printf(name+" start, [PID] %d running...\n", command.Process.Pid)
	Logs(name + "start [PID]=" + strconv.Itoa(command.Process.Pid))
	ioutil.WriteFile(filepath.Join(path, name+".lock"), []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)

}

//停止程序
func Stop(name, path string) {
	pid, err := ioutil.ReadFile(filepath.Join(path, name+".lock"))
	CheckErr(err)
	switch getOs() {
	case 1:
		LinuxKill(name, string(pid))
		break
	case 2:
		WinTaskKill(name, string(pid))
		break
	default:
		break
	}
}

//获取系统类型
func getOs() int {
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

//异常检测
func CheckErr(err error) {
	if err != nil {
		log.Printf("ERROR: #%v ", err)
		Logs(err.Error())
	}

}

//打印日志
func Logs(s string) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fd, _ := os.OpenFile(filepath.Join(dir, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s + "\n")
	defer fd.Close()

}

//Linux kill
func LinuxKill(name, pid string) {
	Exe("kill", "-9", pid)
	println(name + " stop")

}

//win task kill
func WinTaskKill(name, pid string) {
	Exec("TASKKILL", "/PID", pid, "/T", "/F")
	println(name + " stop")
}
