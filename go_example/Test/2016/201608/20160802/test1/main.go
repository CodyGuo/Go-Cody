package main

import (
	"bytes"
	"fmt"
	"io"
)

// const debug = false, when debug is false, out != nil,out == <nil>
const debug = true

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
	fmt.Printf("buf: %s\n", buf)
}

func f(out io.Writer) {
	fmt.Printf("out type: %v\n", out)
	if out != nil {
		fmt.Println("conme in")
		out.Write([]byte("donw."))
	}
}
