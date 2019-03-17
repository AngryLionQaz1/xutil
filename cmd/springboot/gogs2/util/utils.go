package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//获取程序执行目录
//var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

type Project struct {

	//项目路径
	PPath string `yaml:PPath`
	//文件夹名称
	Dir string `yaml:dir`
	//jar包名
	Jar string `yaml:jar`
	//git
	Git string `yaml:git`
	//jar 运行参数
	Arguments string `yaml:arguments`
}

//init
func (p *Project) InitProject() {

	//改变当前工作目录
	os.Chdir(p.PPath)
	//获取文件
	p.git()
	//打包
	p.packageAndRun()

}

//运行jar程序
func (p *Project) Start() {
	//改变当前工作目录
	os.Chdir(filepath.Join(p.PPath, p.Dir))
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
	p.checkErr(err)
	p.wPid(cmd.Process.Pid)
}

//停止jar 程序
func (p *Project) Stop() {
	//改变当前工作目录
	os.Chdir(filepath.Join(p.PPath, p.Dir))
	cmd := exec.Command("kill", "-9", p.rPid())
	err := cmd.Start()
	p.checkErr(err)
}

//运行状态
func (p *Project) Status() {
	str := `
   pid=` + "`" + `ps -ef|grep ` + p.Jar + `|grep -v grep|awk '{print $2}'` + "`" + `
  #如果不存在返回1，存在返回0
  if [ -z "${pid}" ]; then
    echo ">>> ` + p.Jar + ` is not running <<<"
  else
    echo ">>> ` + p.Jar + ` is running PID is ${pid} <<<"
  fi
`
	p.runSh(str)
}

//更新程序
func (p *Project) Update() {
	p.Stop()
	p.InitProject()
	p.Start()
}

//记录pid
func (p *Project) wPid(pid int) {
	fd, _ := os.OpenFile(filepath.Join(p.PPath, p.Dir, p.Dir+".pid"), os.O_RDWR|os.O_CREATE, 0644)
	fd.WriteString(strconv.Itoa(pid))
	defer fd.Close()
}

//读取pid
func (p *Project) rPid() string {
	fd, _ := ioutil.ReadFile(filepath.Join(p.PPath, p.Dir, p.Dir+".pid"))
	return string(fd)
}

//运行sh脚本
func (p *Project) runSh(sh string) {
	os.Chmod(sh, 0777)
	cmd := exec.Command("/bin/bash", "-c", sh)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	p.checkErr(err)
	//执行命令
	if err := cmd.Start(); err != nil {
		p.checkErr(err)
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				p.checkErr(err)
			}
			return
		}
		fmt.Printf("%s\n", string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		p.checkErr(err)
	}
}

//创建文件
func (p *Project) sh(path, str string) {

	fd, _ := os.OpenFile(filepath.Join(p.Dir, path), os.O_RDWR|os.O_CREATE, 0644)
	fd.WriteString(str)
	defer fd.Close()

}

//打包并运行
func (p *Project) packageAndRun() {

	os.Chdir(p.Dir)
	p.exe("mvn", "clean", "package")
	os.Chdir("target")
	p.mv(p.Jar, filepath.Join(p.PPath, p.Dir, p.Jar))

}

//移动文件
func (p *Project) mv(s1, s2 string) {

	err := os.Rename(s1, s2)
	p.checkErr(err)

}

//git 拉取文件
func (p *Project) git() {
	//1，判断文件夹是否存在
	if pthExists(filepath.Join(p.PPath, p.Dir)) {
		os.Chdir(p.Dir)
		ss, _ := os.Getwd()
		fmt.Println(ss)
		p.exe("git", "pull")
	} else {
		p.exe("git", "clone", p.Git)
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

func (p *Project) checkErr(err error) {

	if err != nil {
		log.Println("ERROR: #%v ", err)
		p.logs(err.Error())
	}

}

//打印日志
func (p *Project) logs(s string) {

	fd, _ := os.OpenFile(filepath.Join(p.PPath, p.Dir, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
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
func (p *Project) exe(cm string, args ...string) {
	cmd := exec.Command(cm, args...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	p.checkErr(err)
	//执行命令
	if err := cmd.Start(); err != nil {
		p.checkErr(err)
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
				p.checkErr(err)
			}
			return
		}
		s := string(output)
		fmt.Printf("%s\n", s)
		p.logs(s)
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		p.checkErr(err)
		return
	}
}
