package main

import (
	"log"

	"github.com/cihub/seelog"
	"github.com/fatih/color"
)

func main() {
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")
	log.SetOutput(color.Output)

	log.Println(color.GreenString("Log 日志带颜色吗？"))

	seelog.Default.Info("hello seelog.")
	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")
	seelog.Default.Debug(color.RedString("debug seelog..."))

	seelog.Default.Flush()

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")
}
