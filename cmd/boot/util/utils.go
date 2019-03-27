package util

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/rakyll/statik/fs"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	_ "xutil/cmd/boot/resources"
)

const PACKAGE = "com.snow.golang"
const Name = "golang"

func Run(target, name string) {
	files := GetFiles()
	for _, file := range files {
		output, needHandle, err := ReadFile(file, target, name)
		CheckError(err)
		if needHandle {
			err = WriteToFile(target, filepath.Join(name, file), output)
			CheckError(err)
		}
	}
}

func WriteToFile(target, filePath string, outPut []byte) error {
	ps := PS(target, filePath)
	dir := filepath.Dir(ps)
	CreateDir(dir)
	fmt.Println(ps)
	f, err := os.OpenFile(ps, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func ReadFile(filePath, target, name string) ([]byte, bool, error) {
	//f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	sss, err := fs.New()
	CheckError(err)
	f, err := sss.Open(filePath)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, needHandle, nil
			}
			return nil, needHandle, err
		}
		if ok, _ := regexp.Match(PACKAGE, line); ok {
			reg := regexp.MustCompile(PACKAGE)
			newByte := reg.ReplaceAll(line, []byte(target))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else if ok, _ := regexp.Match(Name, line); ok {
			reg := regexp.MustCompile(Name)
			newByte := reg.ReplaceAll(line, []byte(name))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		}
	}
	return output, needHandle, nil
}

func PS(packages, filePath string) string {
	split := strings.Split(PACKAGE, ".")
	split2 := strings.Split(packages, ".")
	var old bytes.Buffer
	var naw bytes.Buffer
	for _, v := range split {
		old.WriteString(v)
		old.WriteString("\\")
	}
	for _, v := range split2 {
		naw.WriteString(v)
		naw.WriteString("\\")
	}
	return strings.Replace(filePath, old.String(), naw.String(), -1)
}

//获取文件
func GetFiles() []string {
	sss, err := fs.New()
	CheckError(err)
	var files []string
	err = fs.Walk(sss, "/", func(path string, fi os.FileInfo, err error) error {
		CheckError(err)
		files = append(files, path)
		return nil
	})
	CheckError(err)
	sort.Strings(files)
	return files
}

//

//创建目录
func CreateDir(path string) {
	//创建多级目录和设置权限
	os.MkdirAll(path, 0777)
}

//创建文件
func CreateFile(path, content string) {
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666
	fileObj, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err == nil {
		//fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
	}
}

//判断是文件还是文件夹
func CheckDir(f *os.File) bool {
	fs, err := f.Stat()
	CheckError(err)
	return fs.IsDir()
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
