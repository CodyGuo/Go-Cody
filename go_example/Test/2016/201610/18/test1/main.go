package main

import (
	"fmt"
	"io"
	// "os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", `hello world.`)
	stdout, _ := cmd.StdoutPipe()
	// cmd.Stdout = os.Stdout
	// fmt.Println(cmd.Run())
	cmd.Start()

	var buf []byte
	// n, _ := stdout.Read(buf)
	n, _ := io.ReadFull(stdout, buf)

	cmd.Wait()

	fmt.Println(n, string(buf[:n]))
}
