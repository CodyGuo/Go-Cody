package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sum.txt")
	defer file.Close()
	CheckError(err)

	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}

		fmt.Println(string(buf[0:n]))

	}

}

func CheckError(err error) {
	if err != nil {
		fmt.Sprintln(err.Error())
		os.Exit(1)
	}
}
