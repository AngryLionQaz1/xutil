package utils

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

//删除程序
func (p *Project) Delete() {
	p.Stop()
	path := filepath.Join(p.PPath, p.Dir)
	os.RemoveAll(p.PPath)
	logs(path, "delete")
}

//更新程序
func (p *Project) Update() {
	p.Stop()
	p.InitProject()
	p.Start()
	logs(filepath.Join(p.PPath, p.Dir), "update")
}

//停止程序
func (p *Project) Stop() {
	path := filepath.Join(p.PPath, p.Dir)
	//改变当前工作目录
	os.Chdir(path)
	httpPost(p.Actuator)
	logs(path, "stop")
}

//init
func (p *Project) InitProject() {
	//改变当前工作目录
	os.Chdir(p.PPath)
	//获取文件
	git(p)
	//打包
	packageAndRun(p)
	logs(filepath.Join(p.PPath, p.Dir), "init")
}

//运行jar程序
func (p *Project) Start() {
	//改变当前工作目录
	path := filepath.Join(p.PPath, p.Dir)
	os.Chdir(path)
	i := getOs()
	switch i {
	case 1: //Linux
		arguments := []string{"java"}
		for _, v := range strings.Split(p.Arguments, ",") {
			arguments = append(arguments, v)
		}
		ss := []string{"-jar", p.Jar, "/dev/null", "2", ">", "&1", "&"}
		for _, v := range ss {
			arguments = append(arguments, v)
		}
		cmd := exec.Command("nohup", arguments...)
		err := cmd.Start()
		checkErr(err, path)
		break
	case 2: //wind
		arguments := []string{}
		for _, v := range strings.Split(p.Arguments, ",") {
			arguments = append(arguments, v)
		}
		ss := []string{"-jar", p.Jar}
		for _, v := range ss {
			arguments = append(arguments, v)
		}
		cmd := exec.Command("java", arguments...)
		err := cmd.Start()
		checkErr(err, path)
		break
	default:
		break
	}
	logs(path, "start")
}

//http
func httpPost(url string) {
	body_type := "application/json;charset=utf-8"
	http.Post(url, body_type, nil)
}

//打包并运行
func packageAndRun(p *Project) {
	path := filepath.Join(p.PPath, p.Dir)
	path2 := filepath.Join(p.PPath, p.Dir, p.Jar)
	os.Chdir(path)
	exe(path, "mvn", "clean", "package")
	os.Chdir("target")
	mv(path, p.Jar, path2)
}

//移动文件
func mv(path, s1, s2 string) {
	err := os.Rename(s1, s2)
	checkErr(err, path)
}

//git 拉取文件
func git(p *Project) {
	path := filepath.Join(p.PPath, p.Dir)
	//1，判断文件夹是否存在
	if pathExists(path) {
		os.Chdir(path)
		exe(path, "git", "pull")
	} else {
		os.Chdir(p.PPath)
		exe(path, "git", "clone", p.Git)
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

//判断文件或文件夹是否存在
func pathExists(path string) bool {
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
func exe(path, cm string, args ...string) {
	cmd := exec.Command(cm, args...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	checkErr(err, path)
	//执行命令
	if err := cmd.Start(); err != nil {
		checkErr(err, path)
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
			checkErr(err, path)
			return
		}
		s := string(output)
		fmt.Printf("%s\n", s)
		logs(path, s)
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		checkErr(err, path)
		return
	}
}

//打印日志
func logs(path, content string) {
	fd, _ := os.OpenFile(filepath.Join(path, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + content + "\n")
	defer fd.Close()

}

func checkErr(err error, path string) {
	if err != nil {
		logs(path, err.Error())
		log.Println("ERROR:#%v", err)
	}

}
