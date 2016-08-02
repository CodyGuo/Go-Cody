package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

var (
	_SetConsoleTitle uintptr
)

func init() {
	kernel32, loadErr := syscall.LoadLibrary("kernel32.dll")
	if loadErr != nil {
		fmt.Println("loadErr", loadErr)
	}
	defer syscall.FreeLibrary(kernel32)
	_SetConsoleTitle, _ = syscall.GetProcAddress(kernel32, "SetConsoleTitleW")

}

func SetConsoleTitle(title string) int {
	ret, _, callErr := syscall.Syscall(_SetConsoleTitle, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	if callErr != 0 {
		fmt.Println("callErr", callErr)
	}
	return int(ret)
}

func GetFileName(filename string) (string, string) {

	fulleFilename := filepath.Base(filename)

	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fulleFilename)

	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)

	return fulleFilename, filenameOnly
}

func Cprint(str string, num int) {
	fmt.Print("\n")
	for i := 0; i <= num; i++ {
		fmt.Printf(str)
	}
	fmt.Print("\n\n")
}

func main() {
	// -------------------------
	SetConsoleTitle("互普iMan - log解密工具")

	arg_num := len(os.Args)
	fmt.Printf("需要解密的日志文件有 %d 个\n", arg_num-1)

	if arg_num < 2 {
		Cprint("#", 60)
		fmt.Printf("请把需要解密的调试日志文件拖入解密程序中!\n")
	}

	for i := 1; i < arg_num; i++ {
		fulleFilename, filenameOnly := GetFileName(string(os.Args[i]))
		Cprint("#", 60)

		outlog := ".\\logout\\" + filenameOnly + ".rar"
		fmt.Printf("正在解密的第%d个文件: \n"+fulleFilename+"  ----->  "+outlog, i)
		nac_cmd := exec.Command(".\\bin\\openssl", "des3", "-salt", "-d", "-k", "zaq#@!", "-in", os.Args[i], "-out", outlog)
		buf, err := nac_cmd.Output()
		if err != nil {
			Cprint("#", 60)
			fmt.Printf("\n解密失败!请检查bin目录下程序是否完整.包含openssl.exe libeay32.dll ssleay32.dll.\n")
			fmt.Fprintf(os.Stderr, "The command failed to perform: %s %s", err, buf)
		}
	}
	fmt.Printf("\n解密完成,请键入Ctrl+C退出.\n")
	test, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(test)
}
