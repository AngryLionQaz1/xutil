package snow

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"syscall"

	//  "unicode/utf8"
	"unsafe"

	"github.com/axgle/mahonia"
)

type ulong int32
type ulong_ptr uintptr

type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}

type ProcessStruct struct {
	processName string // 进程名称
	processID   int    // 进程id
}

type ProcessStructSlice []ProcessStruct

func (a ProcessStructSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a ProcessStructSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a ProcessStructSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	if strings.Compare(a[j].processName, a[i].processName) < 0 {
		return true
	} else {
		return false
	}
}

func Upayin_process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	_, err := r.Form["callsys"]
	if !err {
		io.WriteString(w, "err")
		return
	}

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	CreateToolhelp32Snapshot := kernel32.NewProc("CreateToolhelp32Snapshot")
	pHandle, _, _ := CreateToolhelp32Snapshot.Call(uintptr(0x2), uintptr(0x0))
	if int(pHandle) == -1 {
		io.WriteString(w, "get process err")
		return
	}
	var data []ProcessStruct
	var buffer bytes.Buffer

	decoder := mahonia.NewDecoder("gbk")

	Process32Next := kernel32.NewProc("Process32Next")
	for {
		var proc PROCESSENTRY32
		proc.dwSize = ulong(unsafe.Sizeof(proc))
		if rt, _, _ := Process32Next.Call(uintptr(pHandle), uintptr(unsafe.Pointer(&proc))); int(rt) == 1 {

			len_szExeFile := 0
			for _, b := range proc.szExeFile {
				if b == 0 {
					break
				}
				len_szExeFile++
			}
			var bytetest []byte = []byte(proc.szExeFile[:len_szExeFile])
			_, newdata, newerr := decoder.Translate(bytetest, true)
			if newerr != nil {
				log.Println(newerr)

			}

			data = append(data, ProcessStruct{
				processName: string(newdata),
				processID:   int(proc.th32ProcessID),
			})

		} else {
			break
		}
	}

	CloseHandle := kernel32.NewProc("CloseHandle")
	_, _, _ = CloseHandle.Call(pHandle)

	sort.Sort(ProcessStructSlice(data))
	for _, v := range data {
		log.Println(v.processName)

		buffer.WriteString("ProcessName : ")
		buffer.WriteString(v.processName)
		buffer.WriteString(" ProcessID : ")
		buffer.WriteString(strconv.Itoa(v.processID))
		buffer.WriteString("\n")
	}

	io.WriteString(w, buffer.String())

}
