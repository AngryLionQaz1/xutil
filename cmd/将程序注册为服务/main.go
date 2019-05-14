package main

import (
	"fmt"
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"xutil/cmd/将程序注册为服务/until"
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
	//程序
	Program string `yaml:program`
	//程序参数
	Args string `yaml:args`
}

func main() {

	program := initProgram()
	config := initService(program)
	s, err := service.New(program, config)
	until.CheckErr(err)
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
	if cmd == "start" {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
	if cmd == "stop" {
		err = program.Stop(s)
		if err != nil {
			log.Fatal(err)
		}

	}

}
func (p *Program) run() {
	path := filepath.Join(p.Path, p.Program)
	// 此处编写具体的服务代码
	until.Start(p.Name, path, p.Args)
}

func (p *Program) Start(s service.Service) error {
	log.Println("开始服务")
	until.Logs("开始服务")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	log.Println("停止服务")
	until.Logs("停止服务")
	until.Stop(p.Name)
	return nil
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

//读取配置文件
func initProgram() *Program {
	i := new(Program)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, initFile))
	until.CheckErr(err)
	err = yaml.Unmarshal(yamlFile, i)
	until.CheckErr(err)
	return i
}
