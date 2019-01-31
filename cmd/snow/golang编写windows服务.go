package snow

import (
	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// 代码写在这儿
	f4, _ := os.OpenFile(`\GoProjects\xutil\asses\4.txt`, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f4.Close()

}

func (p *program) Stop(s service.Service) error {
	return nil
}

/**
* MAIN函数，程序入口
 */

func main() {
	svcConfig := &service.Config{
		Name:        "testServer",    //服务显示名称
		DisplayName: "testServer_my", //服务名称
		Description: "测试服务器",         //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Fatal(err)
	}

	if err != nil {
		logger.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			logger.Println("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {
			s.Uninstall()
			logger.Println("服务卸载成功")
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
