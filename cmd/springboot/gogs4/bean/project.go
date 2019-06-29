package bean

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"xutil/cmd/springboot/gogs4/util"
)

type Project struct {
	//项目路径
	PPath string `yaml:PPath`
	//文件夹名称
	Dir string `yaml:dir`
	//jar包名
	Jar string `yaml:jar`
	//项目停止地址
	Actuator string `yaml:actuator`
	//git
	Git string `yaml:git`
	//jar 运行参数
	Arguments string `yaml:arguments`
	//系统
	Os string `yaml:os`
}

func NewProject(path, dir, jar, actuator, git, arguments string) *Project {
	return &Project{path, dir, jar, actuator, git, arguments, ""}
}

//删除程序
func (p *Project) Delete() {
	p.Stop()
	path := filepath.Join(p.PPath, p.Dir)
	os.RemoveAll(p.PPath)
	util.Logs(path, "delete")
}

//更新程序
func (p *Project) Update() {
	p.Stop()
	p.InitProject()
	p.Start()
	util.Logs(filepath.Join(p.PPath, p.Dir), "update")
}

//停止程序
func (p *Project) Stop() {
	path := filepath.Join(p.PPath, p.Dir)
	//改变当前工作目录
	os.Chdir(path)
	httpPost(p.Actuator)
	util.Logs(path, "stop")
}

//init
func (p *Project) InitProject() {
	//改变当前工作目录
	os.Chdir(p.PPath)
	//获取文件
	util.Git(p.PPath, p.Dir, p.Git)
	//打包
	packageAndRun(p)
	util.Logs(filepath.Join(p.PPath, p.Dir), "init")
}

//运行jar程序
func (p *Project) Start() {
	//改变当前工作目录
	path := filepath.Join(p.PPath, p.Dir)
	os.Chdir(path)
	i := util.GetOs()
	switch i {
	case 1: //Linux
		arguments := []string{"java"}
		for _, v := range strings.Split(p.Arguments, ",") {
			arguments = append(arguments, v)
		}
		ss := []string{"-jar", p.Jar, "/dev/null", "2", ">", "&1", "&"}
		for _, v := range ss {
			arguments = append(arguments, v)
		}
		cmd := exec.Command("nohup", arguments...)
		err := cmd.Start()
		util.CheckErr(err, path)
		break
	case 2: //wind
		arguments := []string{}
		for _, v := range strings.Split(p.Arguments, ",") {
			arguments = append(arguments, v)
		}
		ss := []string{"-jar", p.Jar}
		for _, v := range ss {
			arguments = append(arguments, v)
		}
		cmd := exec.Command("java", arguments...)
		err := cmd.Start()
		util.CheckErr(err, path)
		break
	default:
		break
	}
	util.Logs(path, "start")
}

//http
func httpPost(url string) {
	body_type := "application/json;charset=utf-8"
	http.Post(url, body_type, nil)
}

//打包并运行
func packageAndRun(p *Project) {
	path := filepath.Join(p.PPath, p.Dir)
	path2 := filepath.Join(p.PPath, p.Dir, p.Jar)
	os.Chdir(path)
	util.Exe(path, "mvn", "clean", "package")
	os.Chdir("target")
	util.Mv(path, p.Jar, path2)
}
