package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	debug := flag.Int("debug", 0, "set --debug=1 to debug")
	flag.Parse()
	fmt.Println(*debug)
	os.Exit(1)
}
