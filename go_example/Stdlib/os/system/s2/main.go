package main

import (
	"os"
	"time"
)

func main() {
	name := "./a.txt"
	os.Chmod(name, 0600)

	info, _ := os.Stat(name)
	os.Chtimes(name, time.Now(), info.ModTime())
}
