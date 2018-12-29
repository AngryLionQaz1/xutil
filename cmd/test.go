package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
)

var (
	advapi           = syscall.NewLazyDLL("Advapi32.dll")
	kernel           = syscall.NewLazyDLL("Kernel32.dll")
	user32           = syscall.NewLazyDLL("User32.dll")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

func main() {

	fmt.Printf("X: %d, Y: %d\n", GetSystemMetrics(SM_CXSCREEN), GetSystemMetrics(SM_CYSCREEN))
	tt2()
}

func tt2() {

	fmt.Printf("开机时长:%s\n", GetStartTime())
	fmt.Printf("当前用户:%s\n", GetUserName())
	fmt.Printf("当前系统:%s\n", runtime.GOOS)
	fmt.Printf("系统版本:%s\n", GetSystemVersion())

	fmt.Printf("CPU:\t%s\n", GetCpuInfo())

}

func GetSystemMetrics(nIndex int) int {
	index := uintptr(nIndex)
	ret, _, _ := getSystemMetrics.Call(index)
	return int(ret)
}

//CPU信息
//简单的获取方法fmt.Sprintf("Num:%d Arch:%s\n", runtime.NumCPU(), runtime.GOARCH)
func GetCpuInfo() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	var index = uint32(copy(buffer, syscall.StringToUTF16("Num:")) - 1)
	nums := syscall.StringToUTF16Ptr("NUMBER_OF_PROCESSORS")
	arch := syscall.StringToUTF16Ptr("PROCESSOR_ARCHITECTURE")
	r, err := syscall.GetEnvironmentVariable(nums, &buffer[index], size-index)
	if err != nil {
		return ""
	}
	index += r
	index += uint32(copy(buffer[index:], syscall.StringToUTF16(" Arch:")) - 1)
	r, err = syscall.GetEnvironmentVariable(arch, &buffer[index], size-index)
	if err != nil {
		return syscall.UTF16ToString(buffer[:index])
	}
	index += r
	return syscall.UTF16ToString(buffer[:index+r])
}

//开机时间
func GetStartTime() string {
	GetTickCount := kernel.NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return ""
	}
	ms := time.Duration(r * 1000 * 1000)
	return ms.String()
}

//当前用户名
func GetUserName() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user := syscall.StringToUTF16Ptr("USERNAME")
	domain := syscall.StringToUTF16Ptr("USERDOMAIN")
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
		return ""
	}
	buffer[r] = '@'
	old := r + 1
	if old >= size {
		return syscall.UTF16ToString(buffer[:r])
	}
	r, err = syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
	return syscall.UTF16ToString(buffer[:old+r])
}

//系统版本
func GetSystemVersion() string {
	version, err := syscall.GetVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16)
}

func tt1() {

	s := "$hello,$world $finish"
	fmt.Println(s)

}
