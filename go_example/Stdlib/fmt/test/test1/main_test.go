package main

import (
	"fmt"
	"os"
	"testing"
)

var str = `package main

import (
  "os"

  "github.com/urfave/cli"
)

func main() {
  cli.HelpFlag = cli.BoolFlag{
    Name: "halp, haaaaalp",
    Usage: "HALP",
    EnvVar: "SHOW_HALP,HALPPLZ",
  }

  cli.NewApp().Run(os.Args)
}`

func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < 2; i++ {
		fmt.Print(str)
	}
}

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < 2; i++ {
		os.Stdout.WriteString(str)
	}
}
