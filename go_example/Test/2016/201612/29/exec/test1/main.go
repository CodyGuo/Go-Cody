package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cmd := exec.Command("ping", "-t", "10.10.2.1")
	cmd.Stdout = os.Stdout
	// inW, _ := cmd.StdinPipe()
	cmd.Start()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("退出程序中...", cmd.Process.Pid)
		cmd.Process.Signal(syscall.SIGINT)
		os.Exit(0)
		cmd.Process.Signal(syscall.SIGQUIT)
		cmd.Process.Signal(syscall.SIGTERM)
		// cancel()
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		// Block until a signal is received.
		s := <-c
		fmt.Println("Got signal:", s)
		time.Sleep(5 * time.Second)
		cmd.Process.Signal(os.Interrupt)
	}()

	// cmd.Wait()
}
