package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	b.WriteRune(' ')
	fmt.Fprint(&b, "world.")
	b.WriteByte('\n')
	b.WriteTo(os.Stdout)
}
