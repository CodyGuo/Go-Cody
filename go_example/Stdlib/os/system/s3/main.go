package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.ExpandEnv("HOME = $HOME; GOROOT = $GOROOT"))

	fmt.Println(os.Expand("$GOROOT", func(s string) string {
		return s + " = " + os.Getenv(s)
	}))
}
