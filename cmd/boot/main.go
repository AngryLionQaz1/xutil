package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"xutil/cmd/boot/util"
)

func main() {

	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "boot"
	//程序的用途描述
	app.Usage = "脚手架"
	//程序的版本号
	app.Version = "2.0.0"

	//预置变量
	var projectName string
	var packageName string
	var springBootVersion string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "project, n",
			Value:       "boot",
			Usage:       "项目名",
			Destination: &projectName,
		},
		cli.StringFlag{
			Name:        "package, p",
			Value:       "com.snow.boot",
			Usage:       "包名",
			Destination: &packageName,
		},
		cli.StringFlag{
			Name:        "springboot, s",
			Value:       "2.1.3.RELEASE",
			Usage:       "springboot 版本号",
			Destination: &springBootVersion,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "项目初始化",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName)
				fmt.Println("项目初始化成功")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
