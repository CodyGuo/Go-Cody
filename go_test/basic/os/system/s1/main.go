package main

import (
	"fmt"
	"os"
)

func main() {
	if err := os.Mkdir("a", 0755); os.IsExist(err) {
		fmt.Println("Directory a exists! mkdir Error.", err.Error())
	}

	if file, err := os.Open("../a/test.txt"); os.IsNotExist(err) {
		fmt.Println("File not exists! Cant't Open...")
	} else if os.IsPermission(err) {
		fmt.Println("Permission Denied!")
	} else {
		var buf []byte
		buf = make([]byte, 1024)
		n, _ := file.Read(buf)
		fmt.Println(string(buf[:n]))

		file.Close()
	}
}
