package main

import (
	"fmt"
	"os/exec"
	"strings"
	// "syscall"
)

func main() {
	IP := []string{"10.10.2.215",
		"10.10.2.1",
		"10.10.2.227",
		"10.10.2.162",
		"10.10.3.227",
		"10.10.2.222",
		"10.10.2.116",
		"10.10.2.161"}
	for _, ip := range IP {
		go func() {
			cmd := runCMD(fmt.Sprintf("nmap -O %s", ip))
			result, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(result))
		}()
	}

	for {
	}

}

func runCMD(cmds string) *exec.Cmd {
	cmdList := strings.Split(cmds, " ")
	cmd := exec.Command(cmdList[0], cmdList[1:]...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
