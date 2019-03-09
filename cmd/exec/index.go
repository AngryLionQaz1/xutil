package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

const jar = "C:\\Users\\win\\Desktop\\22s\\xiaoyi.jar"

func main() {

	SysExec("java", "-jar", jar)

}

func exe(cmd string, args ...string) {
	command := exec.Command(cmd, args...)
	buf, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("ERROR: #%v", err)
	}
	fmt.Println(string(buf))
}

/**执行Linux 命令*/
func SysExec(cd string, s ...string) {
	cmd := exec.Command(cd, s...)
	//in:=bytes.NewBuffer(nil)
	//cmd.Stdin=in
	//in.WriteString("y")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		if err != nil {
			log.Fatalf("ERROR: #%v", err)
		}
		return
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		//fmt.Println("Error:The command is err,", err)
		if err != nil {
			log.Fatalf("ERROR: #%v", err)
		}
		return
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {

			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				//fmt.Printf("Error :%s\n", err)
				if err != nil {
					log.Fatalf("ERROR: #%v", err)
				}
			}
			return
		}
		//fmt.Printf("%s\n", string(output))
		fmt.Println(string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		//fmt.Println("wait:", err.Error())
		if err != nil {
			log.Fatalf("ERROR: #%v", err)
		}
		return
	}
}
