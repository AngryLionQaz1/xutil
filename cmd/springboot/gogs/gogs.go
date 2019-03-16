package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"xutil/cmd/springboot/gogs/util"
)

func main() {

	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "gogs"
	//程序的用途描述
	app.Usage = "Linux 运行jar 程序脚本"
	//程序的版本号
	app.Version = "1.0.0"

	app.Commands = []cli.Command{

		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "初始化",
			Action: func(c *cli.Context) error {
				util.InitProject()
				return nil
			},
		},
		{
			Name:    "start",
			Aliases: []string{"st"},
			Usage:   "运行",
			Action: func(c *cli.Context) error {
				util.Start()
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "更新",
			Action: func(c *cli.Context) error {
				util.Update()
				return nil
			},
		},
		{
			Name:    "status",
			Aliases: []string{"ss"},
			Usage:   "状态",
			Action: func(c *cli.Context) error {
				util.Status()
				return nil
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"sp"},
			Usage:   "停止",
			Action: func(c *cli.Context) error {
				util.Stop()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
