package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

const ErrnegativePosition = "bytes.Reader.Seek: negative position"

func main() {
	str := "hello world."
	r := bytes.NewReader([]byte(str))
	n, err := r.Seek(-4, io.SeekEnd)
	checkErr(err)

	fmt.Printf("str --> %s, seek n == %d\n", []byte(str)[n:], n)
}

func checkErr(err error) {
	if err != nil && err.Error() != ErrnegativePosition {
		log.Fatal(err)
	}
}
