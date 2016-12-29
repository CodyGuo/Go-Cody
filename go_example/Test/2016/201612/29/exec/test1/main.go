package main

import (
	"fmt"
	"os/exec"
	// "syscall"
	"time"

	_ "github.com/codyguo/godaemon"
	"github.com/codyguo/win"
)

func main() {
	cmd := exec.Command("ping", "-t", "192.168.1.1")
	// cmd.Stdout = os.Stdout
	// inW, _ := cmd.StdinPipe()
	cmd.Start()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("退出程序中...", cmd.Process.Pid)
		pid := cmd.Process.Pid

		handle := win.OpenProcess(uint32(win.PROCESS_ALL_ACCESS), true, uint32(pid))
		hWnd := win.HWND(handle)
		fmt.Println("进程pid：", pid, "handle: ", hWnd)

		// win.SendMessage(hWnd, win.WM_CLOSE, 0, 0)

		// inW.Write([]byte("{Ctrl}+{C}\r\n"))
		// cmd.Process.Signal(syscall.SIGINT)
		// cmd.Process.Signal(syscall.SIGQUIT)
		// cmd.Process.Signal(syscall.SIGTERM)
		// cmd.Process.Signal(os.Interrupt)
	}()

	cmd.Wait()
}
