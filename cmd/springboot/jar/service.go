package main

import (
	"bufio"
	"fmt"
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Program struct{}
type Init struct {
	//服务名
	Name string `yaml:name`
	//服务器显示
	DisplayName string `yaml:displayName`
	//服务描述
	Description string `yaml:description`
	//jar包名
	Jar string `yaml:jar`
	//停止
	Actuator string `yaml:actuator`
}

var conf *Init

const initFile = "init.yaml"
const (
	Name        = "GoTest"
	DisplayName = "goTest"
	Description = "oserver"
	Path        = "C:\\Users\\pc\\Desktop\\2w"
	Jar         = "xiaoyi.jar"
	Actuator    = "http://127.0.0.1:8086/actuator/shutdown"
)

func main() {

	var serviceConfig = &service.Config{
		Name:        Name,
		DisplayName: DisplayName,
		Description: Description,
	}

	// 构建服务对象
	prog := &Program{}
	s, err := service.New(prog, serviceConfig)
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

//初始换配置文件
func init() {
	i := new(Init)
	yamlFile, err := ioutil.ReadFile(initFile)
	checkErr(err)
	err = yaml.Unmarshal(yamlFile, i)
	checkErr(err)
	conf = i
}

func (p *Program) run() {
	// 此处编写具体的服务代码
	exe("java", "-jar", Jar)
}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	os.Chdir(Path)
	go p.run()
	fd, _ := os.OpenFile("xiaoyi.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	dir, _ := os.Getwd()
	fd.WriteString(dir)
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	body_type := "application/json;charset=utf-8"
	http.Post(Actuator, body_type, nil)
	return nil
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
	}

}
