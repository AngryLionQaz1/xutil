package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"xutil/cmd/将程序注册为服务/until"
)

func main() {

	until.WinTaskKill("tank", "3148")

	//ExecSys("TASKKILL","/PID", "3148","/T", "/F")

}

/**执行基本Linux命令*/
func ExecSys(cd string, ps ...string) {
	cmd := exec.Command(cd, ps...)
	buf, _ := cmd.CombinedOutput()
	b, err := GbkToUtf8(buf)
	fmt.Println(string(b))
	fmt.Println(err)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
