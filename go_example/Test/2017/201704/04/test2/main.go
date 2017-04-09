package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codyguo/logs"
	termbox "github.com/nsf/termbox-go"
)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}

func main() {
	cmd := exec.Command("ping", "www.baidu.com")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd.Wait()

	// var pause bool
	// fmt.Scan(&pause)
	pause()
	logs.Notice("game over.")
}

func pause() {
	fmt.Println("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}
