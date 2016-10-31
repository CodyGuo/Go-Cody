package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func ping(ctx context.Context, ip string) {
	cmd := exec.CommandContext(ctx, "ping", ip, "-t")
	cmd.Stdout = os.Stdin

	go cmd.Run()
	<-ctx.Done()
	fmt.Println("接收到退出消息", cmd.Process.Pid)
	cmd.Process.Kill()
}

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// go ping(ctx, "127.0.0.1")

	ctx, cancel := context.WithCancel(context.Background())
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cancel()
	// }()

	go ping(ctx, "127.0.0.2")
	time.Sleep(3 * time.Second)

	cancel()
	ctx, cancel = context.WithCancel(context.Background())

	go ping(ctx, "127.0.0.3")
	for {
	}

}
