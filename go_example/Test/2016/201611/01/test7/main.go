package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "hello ")
	fmt.Fprint(w, "world.")
	fmt.Fprint(w, "\n")
	w.Flush()
}
