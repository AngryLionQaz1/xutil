package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"xutil/cmd/springCloud分布式框架/cloud/util"
)

func main() {

	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "cloud"
	//程序的用途描述
	app.Usage = "spring cloud alibaba 脚手架"
	//程序的版本号
	app.Version = "1.0.0"

	//预置变量

	var projectName string
	var packageName string
	var springBootVersion string
	var springCloudVersion string
	var springCloudAlibabaVersion string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name, n",
			Value:       "cloud",
			Usage:       "项目名",
			Destination: &projectName,
		},
		cli.StringFlag{
			Name:        "package, p",
			Value:       "com.snow.cloud",
			Usage:       "包名",
			Destination: &packageName,
		},
		cli.StringFlag{
			Name:        "springboot, s",
			Value:       "2.0.6.RELEASE",
			Usage:       "springboot 版本号",
			Destination: &springBootVersion,
		},
		cli.StringFlag{
			Name:        "springCloud, c",
			Value:       "Finchley.SR2",
			Usage:       "springCloud 版本号",
			Destination: &springCloudVersion,
		},
		cli.StringFlag{
			Name:        "springCloudAlibaba, ca",
			Value:       "0.2.1.RELEASE",
			Usage:       "springCloudAlibaba 版本号",
			Destination: &springCloudAlibabaVersion,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "dependencies",
			Aliases: []string{"d"},
			Usage:   "生成项目依赖",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName, util.DEPENDENCIES)
				return nil
			},
		},
		{
			Name:    "commons",
			Aliases: []string{"c"},
			Usage:   "生成项目公共部分",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName, util.COMMONS)
				return nil
			},
		},
		{
			Name:    "auth",
			Aliases: []string{"a"},
			Usage:   "生成认证中心",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName, util.AUTH)
				return nil
			},
		},
		{
			Name:    "gateway",
			Aliases: []string{"g"},
			Usage:   "生成项目网关",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName, util.GATEWAY)
				return nil
			},
		},
		{
			Name:    "module",
			Aliases: []string{"m"},
			Usage:   "生成项目模块",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName, util.GOLANG)
				return nil
			},
		},
		{
			Name:    "all",
			Aliases: []string{"all"},
			Usage:   "生成项目",
			Action: func(c *cli.Context) error {
				util.Run(packageName, projectName, "/")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
