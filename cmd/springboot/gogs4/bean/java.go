package bean

type Java struct {
	Path      string //项目部署路径
	Port      string //项目端口
	Jar       string //jar名称
	Arguments string //运行参数
}

func NewJava(path, port, jar, arguments string) *Java {

	return &Java{Path: path, Port: port, Jar: jar, Arguments: arguments}

}
