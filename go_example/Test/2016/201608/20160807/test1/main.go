package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	_, ok := w.(*os.File)
	fmt.Println(ok)
	_, ok = w.(*bytes.Buffer)
	fmt.Println(ok)

}
