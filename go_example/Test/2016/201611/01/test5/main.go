package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)

	os.Stdout.WriteString("\n")
	var wr bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &wr)
	enc.Write([]byte("Gophers rule!"))
	enc.Close()
	os.Stdout.WriteString(wr.String())
	print(40)

	bufs := bytes.NewBufferString("hello world.")

	fmt.Println(string(bufs.Next(5)))
	print(40)
	fmt.Println(string(bufs.Next(5)))
	print(40)
	// bufs.Reset()
	// bufs.Truncate(1)
	fmt.Println(string(bufs.Next(bufs.Len())))

}

func print(n int) {
	fmt.Println(strings.Repeat("#", n))
}
