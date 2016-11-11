package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	input := []byte("zhangyongfeng@bankofrizhao.com")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	os.Stdout.WriteString("\n")
	encoder.Close()
	pass := []byte("12345678")
	fmt.Println(base64.StdEncoding.EncodeToString(pass))
	fmt.Println('\x00')
}
