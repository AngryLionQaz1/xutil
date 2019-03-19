package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

const (
	ORIGIN = "com.snow.xiaoyi"
	TARGET = "com.snow.boot"
)

func main() {

	files := getFiles(`C:\Users\win\Desktop\xiaoyi`)
	for _, file := range files {
		output, needHandle, err := readFile(file)
		if err != nil {
			panic(err)
		}
		if needHandle {
			err = writeToFile(file, output)
			if err != nil {
				panic(err)
			}
			fmt.Println(file)
		}
	}
	fmt.Println("Done...")

}

func getFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files
}

func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
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

func readFile(filePath string) ([]byte, bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
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
		if ok, _ := regexp.Match(ORIGIN, line); ok {
			reg := regexp.MustCompile(ORIGIN)
			newByte := reg.ReplaceAll(line, []byte(TARGET))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
	return output, needHandle, nil
}
