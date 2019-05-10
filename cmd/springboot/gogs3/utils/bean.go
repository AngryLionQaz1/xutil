package utils

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
