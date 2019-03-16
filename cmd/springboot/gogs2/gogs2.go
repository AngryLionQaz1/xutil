package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"xutil/cmd/springboot/gogs2/util"
)

func main() {

	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "gogs"
	//程序的用途描述
	app.Usage = "Linux 运行jar 程序脚本"
	//程序的版本号
	app.Version = "1.2.0"

	//预置变量
	var projectPath string
	var dir string
	var jar string
	var git string
	var arguments string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "projectPath,p",
			Value:       "/root/test",
			Usage:       "项目路径",
			Destination: &projectPath,
		},
		cli.StringFlag{
			Name:        "dir, d",
			Value:       "file",
			Usage:       "项目文件名",
			Destination: &dir,
		},
		cli.StringFlag{
			Name:        "jar, j",
			Value:       "file-0.0.1-SNAPSHOT.jar",
			Usage:       "jar 包名",
			Destination: &jar,
		},
		cli.StringFlag{
			Name:        "git, g",
			Value:       "root@47.92.213.93:xiaoyiqaz1/file.git",
			Usage:       "git 仓库地址",
			Destination: &git,
		},
		cli.StringFlag{
			Name:        "arguments, a",
			Value:       "-Xms256m,-Xmx512m",
			Usage:       "jar 运行参数",
			Destination: &arguments,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "初始化",
			Action: func(c *cli.Context) error {
				project := util.Project{projectPath, dir, jar, git, arguments}
				fmt.Println(project.Git)
				project.InitProject()
				return nil
			},
		},
		{
			Name:    "start",
			Aliases: []string{"st"},
			Usage:   "运行",
			Action: func(c *cli.Context) error {
				project := util.Project{projectPath, dir, jar, git, arguments}
				project.Start()
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "更新",
			Action: func(c *cli.Context) error {
				project := util.Project{projectPath, dir, jar, git, arguments}
				project.Update()
				return nil
			},
		},
		{
			Name:    "status",
			Aliases: []string{"ss"},
			Usage:   "状态",
			Action: func(c *cli.Context) error {
				project := util.Project{projectPath, dir, jar, git, arguments}
				project.Status()
				return nil
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"sp"},
			Usage:   "停止",
			Action: func(c *cli.Context) error {
				project := util.Project{projectPath, dir, jar, git, arguments}
				project.Stop()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
