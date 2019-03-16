package util

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	initFile = "init.yaml"
)

//获取程序执行目录
var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

var p *Project

type Project struct {

	//文件夹名称
	Dir string `yaml:dir`
	//jar包名
	Jar string `yaml:jar`
	//git
	Git string `yaml:git`
	//jar 运行参数
	Arguments string `yaml:arguments`
}

//读取配置文件
func init() {
	i := new(Project)
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, initFile))
	checkErr(err)
	err = yaml.Unmarshal(yamlFile, i)
	checkErr(err)
	p = i
}

//init
func InitProject() {

	//改变当前工作目录
	os.Chdir(dir)
	//获取文件
	git()
	//打包
	packageAndRun()

}

//运行jar程序
func Start() {
	//改变当前工作目录
	os.Chdir(dir)
	arguments := []string{"java"}
	for _, v := range strings.Split(p.Arguments, ",") {
		arguments = append(arguments, v)
	}
	ss := []string{"-jar", p.Jar, "/dev/null", "2", ">", "&1", "&"}
	for _, v := range ss {
		arguments = append(arguments, v)
	}
	cmd := exec.Command("nohup", arguments...)
	//cmd := exec.Command( "nohup","java","-Xms256m","-Xmx512m","-jar",p.Jar,"/dev/null","2",">","&1","&")
	err := cmd.Start()
	checkErr(err)
	wPid(cmd.Process.Pid)
}

//停止jar 程序
func Stop() {
	//改变当前工作目录
	os.Chdir(dir)
	cmd := exec.Command("kill", "-9", rPid())
	err := cmd.Start()
	checkErr(err)
}

//运行状态
func Status() {
	str := `
   pid=` + "`" + `ps -ef|grep ` + p.Jar + `|grep -v grep|awk '{print $2}'` + "`" + `
  #如果不存在返回1，存在返回0
  if [ -z "${pid}" ]; then
    echo ">>> ` + p.Jar + ` is not running <<<"
  else
    echo ">>> ` + p.Jar + ` is running PID is ${pid} <<<"
  fi
`
	runSh(str)
}

//更新程序
func Update() {
	Stop()
	InitProject()
	Start()
}

//记录pid
func wPid(pid int) {
	fd, _ := os.OpenFile(filepath.Join(dir, p.Dir+".pid"), os.O_RDWR|os.O_CREATE, 0644)
	fmt.Println(pid)
	fd.WriteString(strconv.Itoa(pid))
	defer fd.Close()
}

//读取pid
func rPid() string {
	fd, _ := ioutil.ReadFile(p.Dir + ".pid")
	return string(fd)
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

//创建文件
func sh(path, str string) {

	fd, _ := os.OpenFile(filepath.Join(dir, path), os.O_RDWR|os.O_CREATE, 0644)
	fd.WriteString(str)
	defer fd.Close()

}

//打包并运行
func packageAndRun() {

	os.Chdir(p.Dir)
	exe("mvn", "clean", "package")
	os.Chdir("target")
	mv(p.Jar, filepath.Join(dir, p.Jar))

}

//移动文件
func mv(s1, s2 string) {

	err := os.Rename(s1, s2)
	checkErr(err)

}

//git 拉取文件
func git() {
	//1，判断文件夹是否存在
	if pthExists(p.Dir) {
		exe("git", "pull")
	} else {
		os.Chdir(p.Dir)
		exe("git", "clone", p.Git)
	}
}

//判断文件或文件夹是否存在
func pthExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false

}

func checkErr(err error) {

	if err != nil {
		log.Println("ERROR: #%v ", err)
		logs(err.Error())
	}

}

//打印日志
func logs(s string) {

	fd, _ := os.OpenFile(filepath.Join(dir, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s + "\n")
	defer fd.Close()

}

//创建文件夹
func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(path, 0777)
		// 再修改权限
		os.Chmod(path, 0777)
	}
}

/**执行命令*/
func exe(cm string, args ...string) {
	cmd := exec.Command(cm, args...)
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
		//_, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				checkErr(err)
			}
			return
		}
		s := string(output)
		fmt.Printf("%s\n", s)
		logs(s)
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		checkErr(err)
		return
	}
}
