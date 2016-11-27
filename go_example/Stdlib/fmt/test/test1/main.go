package main

import (
	"fmt"
	"os"
)

func main() {

}

func fmtPrint(str string) {
	fmt.Println(str)
}

func osPrint(str string) {
	os.Stdout.WriteString(str)
}
