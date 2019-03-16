package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (

	//文件夹名称
	Dir = "file"
	//jar包名
	Jar = "file-0.0.1-SNAPSHOT.jar"
	//git
	GitPath = "root@47.92.213.93:xiaoyiqaz1/file.git"
)

//获取程序执行目录
var dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func main() {

	//改变当前工作目录
	os.Chdir(dir)
	//获取文件
	git()
	//打包
	//运行
	packageAndRun()
	//回到初始工作目录
	os.Chdir(dir)
	cmd := exec.Command("nohup", "java", "-Xms256m", "-Xmx512m", "-jar", Jar, "/dev/null", "2", ">", "&1", "&")
	err := cmd.Start()
	checkErr(err)
	wPid(cmd.Process.Pid)
	time.Sleep(60 * time.Second)
	logs(strconv.Itoa(cmd.Process.Pid))
	cm := exec.Command("kill", "-9", rPid())
	cm.Start()

}

//记录pid
func wPid(pid int) {

	fd, _ := os.OpenFile(filepath.Join(dir, Dir+".pid"), os.O_RDWR|os.O_CREATE, 0644)
	fd.WriteString(strconv.Itoa(pid))
	defer fd.Close()

}

//读取pid
func rPid() string {
	fd, _ := ioutil.ReadFile(Dir + ".pid")
	return string(fd)
}

//运行jar程序
func startJar() {
	str := `
JAR=` + Jar + `
PID=` + Dir + `\.pid` + `
isExist(){
  pid=` + "`" + `ps -ef|grep $JAR|grep -v grep|awk '{print $2}'` + "`" + `
  #如果不存在返回1，存在返回0
  if [ -z "${pid}" ]; then
   return 1
  else
    return 0
  fi
}
isExist
if [ $? -eq "0" ]; then
    echo ">>> ${JAR} is already running PID=${pid} <<<"
else
    nohup java -Xms256m -Xmx512m -jar $JAR >/dev/null 2>&1 &
    echo $! > $PID
    echo ">>> start $JAR successed PID=$! <<<"
fi
`
	sh("start.sh", str)
	runSh("start.sh")
}

//停止jar 程序
func stopJar() {

	str := `
JAR=` + Jar + `
PID=` + Dir + `\.pid` + `
isExist(){
  pid=` + "`" + `ps -ef|grep $JAR|grep -v grep|awk '{print $2}'` + "`" + `
  #如果不存在返回1，存在返回0
  if [ -z "${pid}" ]; then
   return 1
  else
    return 0
  fi
}
pidf=$(cat $PID)
echo ">>> api PID = $pidf begin kill $pidf <<<"
kill $pidf
rm -rf $PID
sleep 2
isExist
if [ $? -eq "0" ]; then
    echo ">>> api 2 PID = $pid begin kill -9 $pid  <<<"
    kill -9  $pid
    sleep 2
    echo ">>> $JAR process stopped <<<"
else
    echo ">>> ${JAR} is not running <<<"
fi
`
	sh("stop.sh", str)
	runSh("stop.sh")
}

//运行sh脚本
func runSh(sh string) {
	os.Chmod(sh, 0777)
	cmd := exec.Command("/bin/bash", "-c", "./"+sh)
	err := cmd.Start()
	checkErr(err)
}

//创建文件
func sh(path, str string) {

	fd, _ := os.OpenFile(filepath.Join(dir, path), os.O_RDWR|os.O_CREATE, 0644)
	fd.WriteString(str)
	defer fd.Close()

}

//打包并运行
func packageAndRun() {

	os.Chdir(Dir)
	exe(true, "mvn", "clean", "package")
	os.Chdir("target")
	mv(Jar, filepath.Join(dir, Jar))

}

//移动文件
func mv(s1, s2 string) {

	err := os.Rename(s1, s2)
	checkErr(err)

}

//git 拉取文件
func git() {
	//1，判断文件夹是否存在
	if pthExists(Dir) {
		exe(true, "git", "pull")
	} else {
		exe(true, "git", "clone", GitPath)
	}
}

//判断文件或文件夹是否存在
func pthExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false

}

func checkErr(err error) {

	if err != nil {
		log.Println("ERROR: #%v ", err)
		logs(err.Error())
	}

}

//打印日志
func logs(s string) {

	fd, _ := os.OpenFile(filepath.Join(dir, "logs.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.WriteString(time.Now().Format("2006-01-02 15:04:05") + ":" + s + "\n")
	defer fd.Close()

}

//创建文件夹
func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(path, 0777)
		// 再修改权限
		os.Chmod(path, 0777)
	}
}

/**执行命令*/
func exe(flag bool, cm string, args ...string) *exec.Cmd {
	cmd := exec.Command(cm, args...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	checkErr(err)
	//执行命令
	if err := cmd.Start(); err != nil {
		checkErr(err)
		return nil
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		//_, _, err := outputBuf.ReadLine()
		if err != nil {
			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				checkErr(err)
			}
			return nil
		}
		if flag {
			logs(string(output))
		}
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		checkErr(err)
		return nil
	}
	return cmd
}
