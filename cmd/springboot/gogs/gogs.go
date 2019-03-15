package main

import (
	"bufio"
	"fmt"
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	initFile = "init.yaml"
)

type Program struct {

	//服务名
	Name string `yaml:name`
	//服务器显示
	DisplayName string `yaml:displayName`
	//服务描述
	Description string `yaml:description`
	//路径
	Path string `yaml:path`
	//程序名
	Exe string `yaml:exe`
	//端口
	Port string `yaml:port`
	//参数
	Arguments string `yaml:arguments`
}

func main() {

	program := initProgram()
	config := initService(program)
	s, err := service.New(program, config)
	checkErr(err)
	if len(os.Args) < 2 {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
		return
	}
	cmd := os.Args[1]
	if cmd == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("安装成功")
	}

	if cmd == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("卸载成功")
	}
}

func (p *Program) run() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(dir)
	// 此处编写具体的服务代码

	exe("sh", strings.Split(p.Arguments, ",")...)
}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	logs("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	logs("停止服务")
	exe("kill", "-9", p.Port)
	return nil
}

//读取配置文件
func initProgram() *Program {
	i := new(Program)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, initFile))
	checkErr(err)
	err = yaml.Unmarshal(yamlFile, i)
	checkErr(err)
	return i
}

//初始化服务
func initService(p *Program) *service.Config {

	var serviceConfig = &service.Config{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}
	return serviceConfig
}

/**执行命令*/
func exe(cm string, args ...string) {

	//fd, _ := os.OpenFile("xiaoyi.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
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
		//output, _, err := outputBuf.ReadLine()
		_, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				//fd.Close()
				checkErr(err)
			}
			return
		}
		//fd.WriteString("\n")
		//fd.WriteString(string(output))
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
		logs(err.Error())
	}

}

//打印日志
func logs(s string) {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fd, _ := os.OpenFile(filepath.Join(dir, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s + "\n")
	defer fd.Close()

}
