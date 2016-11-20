package main

import (
	// "log"
	"bufio"
	"io"
	"os"
	"os/exec"
	"runtime"
	// "strings"
	"syscall"
)

import (
	"github.com/axgle/mahonia"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		os.Stdout.WriteString("<iMan tools> ")
		data, _, _ := reader.ReadLine()
		command := string(data)
		cmd := runCmd(command)
		out, _ := cmd.StdoutPipe()
		cmd.Start()
		for {
			result, err := getResult(out)
			if err != nil {
				if err == io.EOF {
					break
				}
				os.Stdout.WriteString("command error: " + command + err.Error())
				break
			}

			os.Stdout.WriteString(result)
		}

		cmd.Wait()
	}

}

func runCmd(cmds string) *exec.Cmd {
	// command := strings.Split(cmds, " ")
	// cmd := exec.Command(command[0], command[1:]...)
	cmd := exec.Command("cmd", "/c", cmds)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd
}

func getResult(out io.ReadCloser) (result string, err error) {
	n := 0
	buf := make([]byte, 1<<16)
	if runtime.GOOS == "windows" {
		gbk := mahonia.NewDecoder("gbk")
		reader := gbk.NewReader(out)
		n, err = reader.Read(buf)
	} else {
		n, err = out.Read(buf)
	}

	return string(buf[:n]), err
}
