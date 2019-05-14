package until

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

/**执行命令*/
func Exe(cm string, args ...string) {
	cmd := exec.Command(cm, args...)
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	CheckErr(err)
	//执行命令
	if err := cmd.Start(); err != nil {
		CheckErr(err)
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
				break
			}
			CheckErr(err)
			return
		}
		s := string(output)
		fmt.Printf("%s\n", s)
		Logs(s)
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		CheckErr(err)
		return
	}
	os.Exit(0)
}

/**执行基本命令*/
func Exec(cd string, ps ...string) {
	cmd := exec.Command(cd, ps...)
	buf, _ := cmd.CombinedOutput()
	b, err := GbkToUtf8(buf)
	fmt.Println(string(b))
	CheckErr(err)
	Logs(string(buf))
}
