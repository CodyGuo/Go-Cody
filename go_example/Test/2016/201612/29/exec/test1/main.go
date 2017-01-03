package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "github.com/codyguo/godaemon"
)

func main() {
	cmd := exec.Command("ping", "10.10.2.162")
	cmd.Stdout = os.Stdout
	// inW, _ := cmd.StdinPipe()
	cmd.Start()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("退出程序中...", cmd.Process.Pid)
		// cmd.Process.Signal(syscall.SIGINT)
		// cmd.Process.Signal(syscall.SIGQUIT)
		// cmd.Process.Signal(syscall.SIGTERM)
		// cancel()
		// c := make(chan os.Signal, 1)
		// signal.Notify(c, os.Interrupt)
		//
		// // Block until a signal is received.
		// s := <-c
		// fmt.Println("Got signal:", s)
		// time.Sleep(5 * time.Second)
		// cmd.Process.Signal(os.Interrupt)

		// cmd.Wait()

		// win.SendMessage(hWnd, win.WM_CLOSE, 0, 0)

		// inW.Write([]byte("{Ctrl}+{C}\r\n"))
		// cmd.Process.Signal(syscall.SIGINT)
		// cmd.Process.Signal(syscall.SIGQUIT)
		// cmd.Process.Signal(syscall.SIGTERM)
		// cmd.Process.Signal(os.Interrupt)
	}()

	cmd.Wait()
}
